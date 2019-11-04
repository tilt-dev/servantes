package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func Random(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm error: %v", err)
	}

	data := r.PostForm["data"]

	// Stop if the client sends us a stop message
	// This is easier than for our testing harness than making them send a signal to the right process
	for _, d := range data {
		if strings.Contains(strings.ToLower(d), "stop") {
			// log.Printf("quitting because we got \"stop\" command in %q", d)
			os.Exit(0)
		}
	}

	if len(data) == 0 {
		fmt.Fprintf(w, "")
		return
	}

	result := data[rand.Intn(len(data))]
	fmt.Fprintf(w, "%s", result)
}

func main() {
	rand.Seed(time.Now().Unix())
	http.HandleFunc("/", Random)
	log.Printf("starting random on :8083")
	http.ListenAndServe(":8083", nil)
}
