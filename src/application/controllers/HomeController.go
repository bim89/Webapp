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

	createTemplate("Index", "home", "layout", usertests[0], res, req)
}

func (HomeController) Create(res http.ResponseWriter, req *http.Request) {
	// renderView(res, req, nil)
	createTemplate("create", "home", "layout", nil, res, req)
}

func (*HomeController) Settings(res http.ResponseWriter, req *http.Request) {
	// renderView(res, req, nil)
	createTemplate("settings", "home", "layout", nil, res, req)
}
