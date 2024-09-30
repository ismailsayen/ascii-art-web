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
			w.WriteHeader(http.StatusInternalServerError)
			asciiartweb.RenderTemplate(w, "./templates/500Page.html", "/")
			return
		}
		w.WriteHeader(http.StatusOK)
	case "/ascii-art":
		if r.Method != "POST" {
		} else if r.Method == "POST" {
			r.ParseForm()
			l := &Data{Text: r.FormValue("text"), Banner: r.FormValue("banner")}
			result := asciiartweb.Ascii(l.Text, l.Banner)
			err := asciiartweb.RenderTemplate(w, "./templates/index.html", result)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				asciiartweb.RenderTemplate(w, "./templates/500Page.html", "/")
				return
			}
			w.WriteHeader(http.StatusOK)
		}
	default:
		err := asciiartweb.RenderTemplate(w, "./templates/404Page.html", "/")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			asciiartweb.RenderTemplate(w, "./templates/500Page.html", "/")
			return
		}
		w.WriteHeader(http.StatusNotFound)

	}
}

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
