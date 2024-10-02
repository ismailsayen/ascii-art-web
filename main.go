package main

import (
	"log"
	"net/http"

	asciiartweb "asciiartweb/Functions"
)

var result string

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		if result != "" {
			asciiartweb.RenderTemplate(w, http.StatusOK, "./templates/index.html", result)
			result = ""
		} else {
			asciiartweb.RenderTemplate(w, http.StatusOK, "./templates/index.html", "")
		}
	default:
		asciiartweb.RenderTemplate(w, http.StatusNotFound, "./templates/404Page.html", "/")
	}
}

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		asciiartweb.RenderTemplate(w, http.StatusBadRequest, "./templates/400Page.html", "/")
	} else if r.Method == "POST" {
		r.ParseForm()
		result = asciiartweb.Ascii(r.FormValue("text"), r.FormValue("banner"))
		http.Redirect(w, r, "/", http.StatusFound)

	}
}

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/ascii-art", HandlerFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
