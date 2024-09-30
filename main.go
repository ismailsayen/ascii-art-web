package main

import (
	"log"
	"net/http"

	asciiartweb "asciiartweb/Functions"
)

type Data struct {
	Text   string
	Banner string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		err := asciiartweb.RenderTemplate(w, "./templates/index.html", "")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			asciiartweb.RenderTemplate(w, "./templates/500Page.html", "/")
			return
		}
	case "/ascii-art":
		if r.Method != "POST" {
		} else if r.Method == "POST" {
			r.ParseForm()
			l := &Data{Text: r.FormValue("text"), Banner: r.FormValue("banner")}
			result := asciiartweb.Ascii(l.Text, l.Banner)
			w.WriteHeader(http.StatusOK)
			err := asciiartweb.RenderTemplate(w, "./templates/index.html", result)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				asciiartweb.RenderTemplate(w, "./templates/500Page.html", "/")
				return
			}
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		asciiartweb.RenderTemplate(w, "./templates/404Page.html", "/")

	}
}

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
