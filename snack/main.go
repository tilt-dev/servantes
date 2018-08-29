package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "The snack of the day is: mochi with green bean")
	})

	http.ListenAndServe(":8083", nil)
}
