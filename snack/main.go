package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		snacks := [...]string{
			"Spam Musubi",
			"Pocky Sticks",
			"Kasugai Gummy",
			"Green Tea Mochi",
			"Shrimp-flavored Chips",
			"Red Bean Rice Cake",
			"Pretz Sticks",
			"Peaches in Agar Jelly",
		}

		rand.Seed(time.Now().Unix())
		s := snacks[rand.Intn(len(snacks))]

		t, err := template.ParseFiles(templatePath("index.tpl"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error parsing template: %v\n", err)
			return
		}

		err = t.Execute(w, s)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error executing template: %v\n", err)
			return
		}
	})

	http.ListenAndServe(":8083", nil)
}

func templatePath(f string) string {
	dir := os.Getenv("TEMPLATE_DIR")
	if dir == "" {
		dir = "snack/web/templates"
	}

	return filepath.Join(dir, f)
}
