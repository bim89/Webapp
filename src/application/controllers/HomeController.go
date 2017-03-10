package controllers

import (
	"net/http"
)

type HomeController struct {
}

func (*HomeController) Index(res http.ResponseWriter, req *http.Request) {
	// renderView(res, req, nil)
	createTemplate("Index", "home", "layout", nil, res, req)
}

func (HomeController) Create(res http.ResponseWriter, req *http.Request) {
	// renderView(res, req, nil)
	createTemplate("create", "home", "layout", nil, res, req)
}

func (*HomeController) Settings(res http.ResponseWriter, req *http.Request) {
	// renderView(res, req, nil)
	createTemplate("settings", "home", "layout", nil, res, req)
}
