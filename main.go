package main

import (
	"log"
	"net/http"

	asciiartweb "asciiartweb/Functions"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		asciiartweb.RenderTemplate(w, http.StatusOK, "./templates/index.html", "")
	default:
		asciiartweb.RenderTemplate(w, http.StatusNotFound, "./templates/404Page.html", "/")
	}
}

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		asciiartweb.RenderTemplate(w, http.StatusBadRequest, "./templates/400Page.html", "/")
	} else if r.Method == "POST" {
		r.ParseForm()
		result := asciiartweb.Ascii(r.FormValue("text"), r.FormValue("banner"))
		asciiartweb.RenderTemplate(w, http.StatusOK, "./templates/index.html", result)
	}
}

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/ascii-art", HandlerFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
