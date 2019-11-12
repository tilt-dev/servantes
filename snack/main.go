package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// The next line creates an error on startup; uncomment it to cause a CrashLoopBackOff
	// log.Fatal("Can't Find Necessary Resource File; dying")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// The next line creates an error on request time; uncomment it to cause an error on request.
		// log.Fatal("NullPointerError trying to service a request")
		snacks := []string{}

		if len(snacks) == 0 {
			snacks = append([]string(nil), defaultSnacks...)
		}

		// Overly-microserviced call to pick a random snack; equivalent to:
		// rand.Seed(time.Now().Unix())
		// s := snacks[rand.Intn(len(snacks))]
		resp, err := http.PostForm(
			fmt.Sprintf("http://%s-random/pick_one", os.Getenv("OWNER")), map[string][]string{
				"data": snacks[:],
			})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error from random server: %v\n", err)
			return
		}

		t, err := template.ParseFiles(templatePath("index.tpl"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error parsing template: %v\n", err)
			return
		}

		defer resp.Body.Close()
		s, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error parsing random response: %v\n", err)
			return
		}

		err = t.Execute(w, string(s))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error executing template: %v\n", err)
			return
		}
	})

	log.Println("Starting Snack Service on :8083")
	http.ListenAndServe(":8083", nil)
}

func templatePath(f string) string {
	dir := os.Getenv("TEMPLATE_DIR")
	return filepath.Join(dir, f)
}

var defaultSnacks = []string{
	"Spam Musubi",
	"Pocky Sticks",
	"Kasugai Gummy",
	"Green Tea Mochi",
	"Shrimp-flavored Chips",
	"Red Bean Rice Cake",
	"Pretz Sticks",
	"Peaches in Agar Jelly",
}
