package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	asciiartweb "asciiartweb/Functions"
)

type Data struct {
	Text   string
	Banner string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/ascii-art":
		t, err := template.ParseFiles("../client/index.html")
		if err != nil {
			fmt.Println("Error: ", err)
		}
		if r.Method == "get" {
			r.ParseForm()
		}
		l := &Data{Text: r.FormValue("text"), Banner: r.FormValue("banner")}
		result := asciiartweb.Ascii(l.Text, l.Banner)
		t.Execute(w, result)
	default:
		w.Write([]byte("hello World"))
	}
}

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8001", nil))
}
