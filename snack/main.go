package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func main() {
	shouldCrash := isCrashFilePresent()

	// The next line creates an error on startup; uncomment it to cause a CrashLoopBackOff
	// log.Fatal("Can't Find Necessary Resource File; dying")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if shouldCrash {
			log.Fatal("NullPointerError trying to service a request")
		}
		// The next line creates an error on request time; uncomment it to cause an error on request.
		// log.Fatal("NullPointerError trying to service a request")
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

	http.HandleFunc("/crash", func(w http.ResponseWriter, r *http.Request) {
		shouldCrash = true
	})

	http.HandleFunc("/uncrash", func(w http.ResponseWriter, r *http.Request) {
		shouldCrash = false
	})

	log.Println("Starting Snack Service on :8083")
	http.ListenAndServe(":8083", nil)
}

func templatePath(f string) string {
	dir := os.Getenv("TEMPLATE_DIR")
	return filepath.Join(dir, f)
}

func isCrashFilePresent() bool {
	_, mainGoPath, _, ok := runtime.Caller(0)
	if !ok {
		return false
	}

	crashFilePath := filepath.Join(filepath.Dir(mainGoPath), "crash.txt")
	_, err := os.Stat(crashFilePath)
	if err != nil {
		return false
	}
	return true
}
