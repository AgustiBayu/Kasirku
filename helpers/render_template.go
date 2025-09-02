package helpers

import (
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, baseDir, tmpl string, data interface{}) {
	tmplParsed := template.Must(template.ParseFiles(baseDir+"/layout.html", baseDir+"/"+tmpl))
	tmplParsed.ExecuteTemplate(w, "layout.html", data)
}
