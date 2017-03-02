package controllers

import (
	"application/models"
	"net/http"
)

type HomeController struct {
}

func (*HomeController) Index(res http.ResponseWriter, req *http.Request) {
	user := models.User{}
	user.Name = "Bjørn-Inge Morstad"
	user.Email = "bimorstad@gmail.com"

	renderView(res, req, user)
}

func (HomeController) Create(res http.ResponseWriter, req *http.Request) {
	// renderView(res, req, nil)
	createTemplate("create", "home", "layout", nil, res, req)
}

func (*HomeController) Settings(res http.ResponseWriter, req *http.Request) {
	renderView(res, req, nil)
}
