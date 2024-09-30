package asciiartweb

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, status int, tmpl string, data string) error {
	w.WriteHeader(status)
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return err
	}

	err = t.Execute(w, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return err
	}

	return nil
}
