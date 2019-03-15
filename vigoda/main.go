package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	go func() {
		for {
			time.Sleep(2 * time.Second)
			log.Printf("Server status: All good")
		}
	}()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time := "10:00am"

		fmt.Fprintf(w, time)
		log.Printf("Yep, still dead.")
	})

	log.Println("Starting Vigoda Health Check Service on :8081")
	http.ListenAndServe(":8081", nil)
}

func templatePath(f string) string {
	dir := os.Getenv("TEMPLATE_DIR")
	return filepath.Join(dir, f)
}
