package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"regexp"
	"runtime"
)

const defaultLayout = "layout"

func templatesDir(directory string) string {
	return fmt.Sprintf("src/application/templates/%s", directory)
}

func renderView(res http.ResponseWriter, req *http.Request, data interface{}) {

	pc := make([]uintptr, 100) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	var folder string = matchString(`\(\*(.*?)Controller\)`, f.Name())
	var file string = matchString(`Controller\).(.*)`, f.Name())

	createTemplate(file, folder, defaultLayout, data, res, req)
}

func renderViewWithLayout(layout string, data interface{}, res http.ResponseWriter, req *http.Request) {
	pc := make([]uintptr, 100)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	var folder string = matchString(`\(\*(.*?)Controller\)`, f.Name())
	var file string = matchString(`Controller\).(.*)`, f.Name())

	createTemplate(file, folder, layout, data, res, req)
}

func renderViewWithFile(file string, data interface{}, res http.ResponseWriter, req *http.Request) {
	pc := make([]uintptr, 100)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	var folder string = matchString(`\(\*(.*?)Controller\)`, f.Name())

	createTemplate(file, folder, defaultLayout, data, res, req)
}

func createTemplate(file string, folder string, layout string, data interface{}, res http.ResponseWriter, req *http.Request) {
	lp := path.Join(templatesDir(layout), fmt.Sprintf("%s.html", layout))
	fp := path.Join(templatesDir(folder), fmt.Sprintf("%s.html", file))

	tmpl, err := template.ParseFiles(fp, lp)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(res, layout, data); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func matchString(r string, s string) string {
	reg := regexp.MustCompile(r)
	match := reg.FindStringSubmatch(s)
	println(match[1])

	return match[1]
}
