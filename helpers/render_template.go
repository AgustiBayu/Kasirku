package helpers

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	files := []string{
		"templates/layouts/main.html",
		tmpl,
	}

	tmplParsed, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Error parsing template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmplParsed.ExecuteTemplate(w, "main.html", data)
	if err != nil {
		println("Error executing template:", err.Error())
	}
}