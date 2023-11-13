package templates

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func MustParseTemplate(filename string) *template.Template {
	path := filepath.Join("templates", filename)
	tpl, err := template.ParseFiles("templates/layout.html", path)
	if err != nil {
		panic(err)
	}
	return tpl
}

func RenderTemplate(w http.ResponseWriter, tpl *template.Template, data any) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	err := tpl.Execute(w, data)
	if err != nil {
		log.Printf("Render Template: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
	}
}

var IndexTemplate = MustParseTemplate("index.html")
var ShowProductTemplate = MustParseTemplate("products/show.html")
