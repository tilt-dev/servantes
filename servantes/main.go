package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/duration"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var Services = []string{
	"fortune",
	"vigoda",
	"doggos",
}

var ProxyMap = make(map[string]*httputil.ReverseProxy, 0)
var Client *kubernetes.Clientset

func main() {
	client, err := createK8sClient()
	if err != nil {
		log.Printf("Error initializing k8s client: %v\n")
	}
	Client = client

	for _, s := range Services {
		ProxyMap[s] = newProxy(s)
	}

	r := mux.NewRouter()
	r.PathPrefix("/s/{service}").Handler(http.HandlerFunc(handleService))
	r.HandleFunc("/", handleIndex)
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(templatePath("index.tpl"))
	if err != nil {
		http.Error(w, fmt.Sprintf("error parsing template: %v\n", err),
			http.StatusInternalServerError)
		return
	}

	services := listServices()
	err = t.Execute(w, templateData{Services: services})
	if err != nil {
		http.Error(w, fmt.Sprintf("error executing template: %v\n", err),
			http.StatusInternalServerError)
		return
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

// List all services in our hard-coded list by querying the k8s api.
func listServicesFromK8sAPI() (map[string]serviceData, error) {
	serviceMap := make(map[string]serviceData, 0)
	if Client == nil {
		return serviceMap, nil
	}

	podList, err := Client.CoreV1().Pods("default").List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, pod := range podList.Items {
		app := pod.ObjectMeta.Labels["app"]
		phase := pod.Status.Phase
		bestStartTime := bestStartTime(pod)

		_, isServantes := ProxyMap[app]
		if !isServantes {
			continue
		}

		existingService := serviceMap[app]
		if existingService.StartTime.After(bestStartTime) {
			continue
		}

		data := serviceData{
			Name:      app,
			Phase:     string(phase),
			StartTime: bestStartTime,
		}
		serviceMap[app] = data
	}

	return serviceMap, nil
}

// We want to get the time that best reflects the current age
// of the container running the service.
func bestStartTime(pod v1.Pod) time.Time {
	podStartTime := pod.Status.StartTime.Time
	bestStartTime := podStartTime

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
func listServices() []serviceData {
	serviceMap, err := listServicesFromK8sAPI()
	if err != nil {
		log.Printf("Error fetching pods: %v\n", err)
		serviceMap = make(map[string]serviceData, 0)
	}

	result := make([]serviceData, 0)
	for _, service := range Services {
		data, found := serviceMap[service]
		if found {
			result = append(result, data)
		} else {
			result = append(result, serviceData{
				Name:  service,
				Phase: "Unknown",
			})
		}
	}
	return result
}

func templatePath(f string) string {
	dir := os.Getenv("TEMPLATE_DIR")
	if dir == "" {
		dir = "servantes/web/templates"
	}
	return filepath.Join(dir, f)
}

func handleService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := vars["service"]
	proxy, isValid := ProxyMap[service]
	if !isValid {
		http.Error(w, fmt.Sprintf("Service %q not found\nAvailable services: %v\n", service, Services),
			http.StatusNotFound)
		return
	}

	proxy.ServeHTTP(w, r)
}

func newProxy(service string) *httputil.ReverseProxy {
	prefix := fmt.Sprintf("/s/%s", service)
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = service
		req.URL.Path = strings.Replace(req.URL.Path, prefix, "", 1)
		if req.URL.Path == "" {
			req.URL.Path = "/"
		}

	}
	return &httputil.ReverseProxy{Director: director}
}

type templateData struct {
	Services []serviceData
}

type serviceData struct {
	Name      string
	Phase     string
	StartTime time.Time
}

func (d serviceData) HumanAge() string {
	return duration.ShortHumanDuration(time.Since(d.StartTime))
}
