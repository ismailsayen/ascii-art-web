package asciiartweb

import (
	"fmt"
	"net/http"
	"text/template"

	asciiartweb "asciiartweb/ascii-art"
)

type Data struct {
	Text   string
	Banner string
}

func ConnectTemplate(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println(result)
}
