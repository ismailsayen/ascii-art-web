package asciiartweb

import (
	"fmt"
	"net/http"
	"text/template"
)

func ConnectTemplate(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../client/index.html")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	t.Execute(w, nil)
	if r.Method == "get" {
		r.ParseForm()
	}
	
	fmt.Fprint(w,"-Text: " ,r.FormValue("text"), " -banner: ", r.FormValue("banner"))


}
