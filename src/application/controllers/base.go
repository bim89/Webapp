package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"regexp"
	"runtime"
	"path/filepath"
	"os"
	"log"
)

const defaultLayout = "layout"

func templatesDir(directory string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	return fmt.Sprintf("src/application/templates/%s", directory)
}

func layoutDir(directory string) string {
	return fmt.Sprintf("src/application/templates/layout/")
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

	lp := path.Join(layoutDir(layout), fmt.Sprintf("%s.html", layout))
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

func serveFile(file string, folder string, data interface{}, res http.ResponseWriter, req *http.Request) {
	fp := path.Join(templatesDir(folder), fmt.Sprintf("%s.html", file))
	t, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := t.ExecuteTemplate(res, fp, nil); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func matchString(r string, s string) string {
	reg := regexp.MustCompile(r)
	match := reg.FindStringSubmatch(s)
	println(match[1])

	return match[1]
}
