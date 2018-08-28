package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Abe Vigoda is still dead as of 9:54AM.")
	})

	http.ListenAndServe(":8081", nil)
}
