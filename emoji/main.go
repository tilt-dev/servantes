package main

import (
	"fmt"
	"github.com/windmilleng/servantes/emoji/pkg/emoji"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type TemplateArgs struct {
	EmojiString string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		str := ""
		numEmoji := rnd.Intn(6) + 1
		for i := 0; i < numEmoji; i++ {
			str += string(emoji.RandomEmoji(rnd))
		}
		t, err := template.ParseFiles(templatePath("index.tpl"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error parsing template: %v\n", err)
			return
		}

		templateArgs := TemplateArgs{str}

		err = t.Execute(w, templateArgs)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error executing template: %v\n", err)
			return
		}

	})

	http.ListenAndServe(":8081", nil)
}

func templatePath(f string) string {
	dir := os.Getenv("TEMPLATE_DIR")
	return filepath.Join(dir, f)
}
