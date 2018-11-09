package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var morningSnacks = []string{
	"Spam Musubi",
	"Kasugai Gummy",
	"Green Tea Mochi",
	"Sous Vide Eggs",
}

var afternoonSnacks = []string{
	"Shrimp-flavored Chips",
	"Red Bean Rice Cake",
	"Pretz Sticks",
	"Peaches in Agar Jelly",
	"Pocky Sticks",
}

var afternoon = flag.Bool("afternoon", false, "Uses the afternoon snack list instead of the morning snack list")

type SnackData struct {
	Type string
	Name string
}

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		snacks := morningSnacks
		if *afternoon {
			snacks = afternoonSnacks
		}

		rand.Seed(time.Now().Unix())
		s := snacks[rand.Intn(len(snacks))]

		t, err := template.ParseFiles(templatePath("index.tpl"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error parsing template: %v\n", err)
			return
		}

		err = t.Execute(w, NewSnackData(s))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error executing template: %v\n", err)
			return
		}
	})

	log.Println("Starting Snack Service on :8083")
	if *afternoon {
		log.Println("Serving afternoon snacks because --afternoon was set!")
	} else {
		log.Println("Serving morning snacks!")
	}
	http.ListenAndServe(":8083", nil)
}

func NewSnackData(s string) SnackData {
	t := "morning"
	if *afternoon {
		t = "afternoon"
	}
	return SnackData{Type: t, Name: s}
}

func templatePath(f string) string {
	dir := os.Getenv("TEMPLATE_DIR")
	return filepath.Join(dir, f)
}
