package asciiartweb

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, status int, tmpl string, data string) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, "We're sorry, but something went wrong on our end. Please try again later.", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}
}
