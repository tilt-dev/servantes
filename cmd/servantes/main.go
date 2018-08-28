package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

var Services = []string{
	"fortune",
	"vigoda",
}

var ProxyMap = make(map[string]*httputil.ReverseProxy, 0)

func main() {
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

	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("error executing template: %v\n", err),
			http.StatusInternalServerError)
		return
	}
}

func templatePath(f string) string {
	dir := os.Getenv("TEMPLATE_DIR")
	if dir == "" {
		dir = "web/servantes/templates"
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
