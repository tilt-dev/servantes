package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time := "10:00am"

		t, err := template.ParseFiles(templatePath("index.tpl"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error parsing template: %v\n", err)
			return
		}

		err = t.Execute(w, time)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error executing template: %v\n", err)
			return
		}
		log.Printf("Yep, still dead.")
	})

	log.Println("Starting Vigoda Health Check Service on :8081")
	http.ListenAndServe(":8081", nil)
}

func templatePath(f string) string {
	dir := os.Getenv("TEMPLATE_DIR")
	return filepath.Join(dir, f)
}
