package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/duration"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type ServiceSpec struct {
	Name string
	Port int
}

var ServiceSpecs = []ServiceSpec{
	{"fortune", 9004},
	{"vigoda", 9001},
	{"snack", 9002},
	{"doggos", 9003},
	{"hypothesizer", 9005},
	{"spoonerisms", 9006},
	{"emoji", 9007},
}

var ServiceSpecMap = make(map[string]ServiceSpec, len(ServiceSpecs))

var Client *kubernetes.Clientset

var serviceOwner = flag.String("owner", "", "If specified, servantes will only use services that have an `owner` label with the given value")

func main() {
	flag.Parse()

	client, err := createK8sClient()
	if err != nil {
		log.Printf("Error initializing k8s client: %v\n", err)
	}
	Client = client

	for _, s := range ServiceSpecs {
		ServiceSpecMap[s.Name] = s
	}

	r := mux.NewRouter()
	r.HandleFunc("/", handleIndex(*serviceOwner))
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}

func handleIndex(serviceOwner string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(templatePath("index.tpl"))
		if err != nil {
			http.Error(w, fmt.Sprintf("error parsing template: %v\n", err),
				http.StatusInternalServerError)
			return
		}

		services := listServices(serviceOwner)
		err = t.Execute(w, templateData{Services: services})
		if err != nil {
			http.Error(w, fmt.Sprintf("error executing template: %v\n", err),
				http.StatusInternalServerError)
			return
		}
	}
}

func createK8sClient() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// copied from https://github.com/kubernetes/kubernetes/blob/aedeccda9562b9effe026bb02c8d3c539fc7bb77/pkg/kubectl/resource_printer.go#L692-L764
// to match the status column of `kubectl get pods`
func getStatus(pod v1.Pod) (string, int) {
	restarts := 0

	reason := string(pod.Status.Phase)
	if pod.Status.Reason != "" {
		reason = pod.Status.Reason
	}

	initializing := false
	for i := range pod.Status.InitContainerStatuses {
		container := pod.Status.InitContainerStatuses[i]
		restarts += int(container.RestartCount)
		switch {
		case container.State.Terminated != nil && container.State.Terminated.ExitCode == 0:
			continue
		case container.State.Terminated != nil:
			// initialization is failed
			if len(container.State.Terminated.Reason) == 0 {
				if container.State.Terminated.Signal != 0 {
					reason = fmt.Sprintf("Init:Signal:%d", container.State.Terminated.Signal)
				} else {
					reason = fmt.Sprintf("Init:ExitCode:%d", container.State.Terminated.ExitCode)
				}
			} else {
				reason = "Init:" + container.State.Terminated.Reason
			}
			initializing = true
		case container.State.Waiting != nil && len(container.State.Waiting.Reason) > 0 && container.State.Waiting.Reason != "PodInitializing":
			reason = "Init:" + container.State.Waiting.Reason
			initializing = true
		default:
			reason = fmt.Sprintf("Init:%d/%d", i, len(pod.Spec.InitContainers))
			initializing = true
		}
		break
	}
	if !initializing {
		restarts = 0
		for i := len(pod.Status.ContainerStatuses) - 1; i >= 0; i-- {
			container := pod.Status.ContainerStatuses[i]

			restarts += int(container.RestartCount)
			if container.State.Waiting != nil && container.State.Waiting.Reason != "" {
				reason = container.State.Waiting.Reason
			} else if container.State.Terminated != nil && container.State.Terminated.Reason != "" {
				reason = container.State.Terminated.Reason
			} else if container.State.Terminated != nil && container.State.Terminated.Reason == "" {
				if container.State.Terminated.Signal != 0 {
					reason = fmt.Sprintf("Signal:%d", container.State.Terminated.Signal)
				} else {
					reason = fmt.Sprintf("ExitCode:%d", container.State.Terminated.ExitCode)
				}
			}
		}
	}

	return reason, restarts
}

// List all services in our hard-coded list by querying the k8s api.
func listServicesFromK8sAPI(serviceOwner string) (map[string]serviceData, error) {
	serviceMap := make(map[string]serviceData, 0)
	if Client == nil {
		return serviceMap, nil
	}

	podList, err := Client.CoreV1().Pods("default").List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, pod := range podList.Items {
		if len(serviceOwner) > 0 && pod.ObjectMeta.Labels["owner"] != serviceOwner {
			continue
		}

		app := pod.ObjectMeta.Labels["app"]
		spec, isServantes := ServiceSpecMap[app]
		if !isServantes {
			continue
		}

		status, restartCount := getStatus(pod)
		bestStartTime := bestStartTime(pod)

		existingService := serviceMap[app]
		if existingService.StartTime.After(bestStartTime) {
			continue
		}

		data := serviceData{
			Name:         spec.Name,
			Port:         spec.Port,
			Status:       status,
			RestartCount: restartCount,
			StartTime:    bestStartTime,
		}

		serviceMap[app] = data
	}

	return serviceMap, nil
}

// We want to get the time that best reflects the current age
// of the container running the service.
func bestStartTime(pod v1.Pod) time.Time {
	var bestStartTime time.Time

	podStartTime := pod.Status.StartTime
	if podStartTime != nil {
		bestStartTime = podStartTime.Time
	}

	for _, cStatus := range pod.Status.ContainerStatuses {
		state := cStatus.State
		if state.Running != nil {
			cStartTime := state.Running.StartedAt.Time
			if cStartTime.After(bestStartTime) {
				bestStartTime = cStartTime
			}
		}
	}
	return bestStartTime
}

// Format all services as serviceData objects.
// Fail back to the hard-coded list if we don't have access to the k8s api.
func listServices(serviceOwner string) []serviceData {
	serviceMap, err := listServicesFromK8sAPI(serviceOwner)
	if err != nil {
		log.Printf("Error fetching pods: %v\n", err)
		serviceMap = make(map[string]serviceData, 0)
	}

	result := make([]serviceData, 0)
	for _, service := range ServiceSpecs {
		data, found := serviceMap[service.Name]
		if found {
			result = append(result, data)
		} else {
			result = append(result, serviceData{
				Name:   service.Name,
				Status: "Unknown",
			})
		}
	}
	return result
}

func templatePath(f string) string {
	dir := os.Getenv("TEMPLATE_DIR")
	if dir == "" {
		dir = "fe/web/templates"
	}
	return filepath.Join(dir, f)
}

type templateData struct {
	Services []serviceData
}

type serviceData struct {
	Name         string
	Port         int
	Status       string
	RestartCount int
	StartTime    time.Time
}

func (d serviceData) HumanAge() string {
	return duration.ShortHumanDuration(time.Since(d.StartTime))
}
