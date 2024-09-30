package main

import (
	"log"
	"net/http"

	asciiartweb "asciiartweb/Functions"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		w.WriteHeader(http.StatusOK)
		err := asciiartweb.RenderTemplate(w, http.StatusOK, "./templates/index.html", "")
		if err != nil {
			http.Error(w, "We're sorry, but something went wrong on our end. Please try again later.", http.StatusInternalServerError)
			return
		}
	case "/ascii-art":
		if r.Method != "POST" {
			err := asciiartweb.RenderTemplate(w, http.StatusBadRequest, "./templates/400Page.html", "/")
			if err != nil {
				http.Error(w, "We're sorry, but something went wrong on our end. Please try again later.", http.StatusInternalServerError)
				return
			}
		} else if r.Method == "POST" {
			r.ParseForm()
			result := asciiartweb.Ascii(r.FormValue("text"), r.FormValue("banner"))
			err := asciiartweb.RenderTemplate(w, http.StatusOK, "./templates/index.html", result)
			if err != nil {
			}
		}
	default:
		err := asciiartweb.RenderTemplate(w, http.StatusNotFound, "./templates/404Page.html", "/")
		if err != nil {
			http.Error(w, "We're sorry, but something went wrong on our end. Please try again later.", http.StatusInternalServerError)
			return
		}

	}
}

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
