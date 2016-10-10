package controllers

import (
	"fmt"
	"net/http"
	"path"
	"html/template"
)

const defaultLayout = "layout"

func templatesDir(directory string) string {
	return fmt.Sprintf("src/application/templates/%s", directory)
}

func renderView(layout string, tmplDir string, tmplFile string, data interface{},res http.ResponseWriter, req *http.Request) {
	lp := path.Join(templatesDir("layout"), fmt.Sprintf("%s.html", layout))
	fp := path.Join(templatesDir(tmplDir), fmt.Sprintf("%s.html", tmplFile))


	tmpl, err := template.ParseFiles(fp, lp)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(res, layout, data); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
