package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/windmilleng/servantes/api/fortune"
)

func main() {
	f := fortune.Fortune{Text: "you will have a nice day"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(templatePath("index.tpl"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error parsing template: %v\n", err)
			return
		}

		err = t.Execute(w, f.Text)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error executing template: %v\n", err)
			return
		}
	})

	http.ListenAndServe(":8082", nil)
}

func templatePath(f string) string {
	dir := os.Getenv("TEMPLATE_DIR")
	if dir == "" {
		dir = "web/fortune/templates"
	}

	return filepath.Join(dir, f)
}
