package main

import (
	"log"
	"net/http"

	asciiartweb "asciiartweb/template"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/ascii-art":
		asciiartweb.ConnectTemplate(w, r)
	default:
		w.Write([]byte("hello World"))
	}
}

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8001", nil))
}
