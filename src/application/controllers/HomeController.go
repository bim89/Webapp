package controllers

import (
	"net/http"
	"application/models"
)

type HomeController struct {
}

func (*HomeController) Index(res http.ResponseWriter, req *http.Request) {
	// renderView(res, req, nil)
	ut := models.UserTest{}
	usertests := ut.FindAll()

	createTemplate("index", "home", "layout", usertests, res, req)

}

func (HomeController) Create(res http.ResponseWriter, req *http.Request) {
	// renderView(res, req, nil)
	createTemplate("create", "home", "layout", nil, res, req)
}

func (HomeController) Show(res http.ResponseWriter, req *http.Request) {
	// renderView(res, req, nil
	id := req.FormValue("t")
	// ut := models.UserTest{}

	createTemplate("show", "home", "layout", id, res, req)
}

func (*HomeController) Settings(res http.ResponseWriter, req *http.Request) {
	// renderView(res, req, nil)
	createTemplate("settings", "home", "layout", nil, res, req)
}

func (*HomeController) Login(res http.ResponseWriter, req *http.Request) {
	// renderView(res, req, nil)
	// createTemplate("login", "home", "", nil, res, req)
	createTemplate("login_form", "home", "login", nil, res, req)
}
