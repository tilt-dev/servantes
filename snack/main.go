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
		snacks := [...]string{
			"Snack 1",
		}

		resp, err := http.PostForm(
			fmt.Sprintf("http://%s-random", os.Getenv("OWNER")), map[string][]string{
				"data": snacks[:],
			})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error from random server: %v\n", err)
			return
		}
		// rand.Seed(time.Now().Unix())
		// s := snacks[rand.Intn(len(snacks))]

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
