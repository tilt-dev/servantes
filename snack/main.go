package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
		fmt.Fprintf(w, "Your next snack: %s", snacks[rand.Intn(len(snacks))])

	})

	http.ListenAndServe(":8083", nil)
}
