package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/windmilleng/servantes/emoji/pkg/emoji"
)

type TemplateArgs struct {
	EmojiRows []string
}

var dimension = flag.Int("dimension", 1, "number of rows of emoji to print")

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		rows := []string{}
		numEmoji := rnd.Intn(6) + 1
		for d := 0; d < *dimension; d++ {
			str := ""
			for i := 0; i < numEmoji; i++ {
				str += string(emoji.RandomEmoji(rnd))
			}
			rows = append(rows, str)
		}
		t, err := template.ParseFiles(templatePath("index.tpl"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error parsing template: %v\n", err)
			return
		}

		templateArgs := TemplateArgs{EmojiRows: rows}

		err = t.Execute(w, templateArgs)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error executing template: %v\n", err)
			return
		}

	})

	log.Println("Starting Emoji Service on :8081")
	log.Printf("Printing %d rows of emoji\n", *dimension)
	http.ListenAndServe(":8081", nil)
}

func templatePath(f string) string {
	dir := os.Getenv("TEMPLATE_DIR")
	return filepath.Join(dir, f)
}
