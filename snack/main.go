package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// The next line creates an error on startup; uncomment it to cause a CrashLoopBackOff
	// log.Fatal("Can't Find Necessary Resource File; dying")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
			"hello world",
		}

		rand.Seed(time.Now().Unix())
		s := snacks[rand.Intn(len(snacks))]

		_, err := fmt.Fprintf(w, s)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error writing response: %v\n", err)
		}
	})

	log.Println("Starting Snack Service on :8083")
	http.ListenAndServe(":8083", nil)
}

func templatePath(f string) string {
	dir := os.Getenv("TEMPLATE_DIR")
	return filepath.Join(dir, f)
}
