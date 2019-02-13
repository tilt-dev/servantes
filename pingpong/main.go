package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Got a ping!")
		fmt.Fprintf(w, "pong!\n")
	})

	log.Println("Starting Ping/Pong Service on :8081")
	http.ListenAndServe(":8081", nil)
}
