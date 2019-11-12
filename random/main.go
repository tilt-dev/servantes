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

const quitThreshold = 5

var quitCount = 0

func Random(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm error: %v", err)
	}

	data := r.PostForm["data"]

	// Stop if the client sends us N quit messages
	// This is easier than for our testing harness than making them send a signal to the right process
	for _, d := range data {
		if strings.Contains(strings.ToLower(d), "quit") || strings.Contains(strings.ToLower(d), "stop") {
			// log.Printf("got quit command %q", d)
			quitCount++
			if quitCount > quitThreshold {
				log.Fatalf("commanded to quit")
				os.Exit(0)
			}
		}
	}

	if len(data) == 0 {
		fmt.Fprintf(w, "")
		return
	}

	// test

	// test 3
	// test 4
	// test 2
	// test 5
	// test 6
	// test 7
	// test 8
	// test 9
	result := data[rand.Intn(len(data))]
	fmt.Fprintf(w, "%s", result)
}

func main() {
	rand.Seed(time.Now().Unix())
	http.HandleFunc("/pick_one", Random)
	log.Printf("starting random on :8083")
	http.ListenAndServe(":8083", nil)
}
