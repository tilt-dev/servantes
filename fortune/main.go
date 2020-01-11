package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/windmilleng/servantes/fortune/api"
)

func main() {
	f := api.Fortune{
		Text: "you will have a nice day",
		Secret: os.Getenv("THE_SECRET"),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(templatePath("index.tpl"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error parsing template: %v\n", err)
			return
		}

		err = t.Execute(w, f)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error executing template: %v\n", err)
			return
		}
	})

	log.Println("Starting Fortune Service on :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func templatePath(f string) string {
	dir := os.Getenv("TEMPLATE_DIR")
	if dir == "" {
		dir = "fortune/web/templates"
	}

	return filepath.Join(dir, f)
}
