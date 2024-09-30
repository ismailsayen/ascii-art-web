package asciiartweb

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data string) error {
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
