package controllers

import (
	"net/http"
	"application/models"
)

type HomeController struct {
}

type User struct {
	Title string
	Email string
}

func (*HomeController) Index(res http.ResponseWriter, req *http.Request) {
	// renderView(res, req, nil)
	ut := models.UserTest{}
	usertests := ut.FindAll()

	createTemplate("Index", "home", "layout", usertests, res, req)
}

func (HomeController) Create(res http.ResponseWriter, req *http.Request) {
	// renderView(res, req, nil)
	createTemplate("create", "home", "layout", nil, res, req)
}

func (*HomeController) Settings(res http.ResponseWriter, req *http.Request) {
	// renderView(res, req, nil)
	createTemplate("settings", "home", "layout", nil, res, req)
}
