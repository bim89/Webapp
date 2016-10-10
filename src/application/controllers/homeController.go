package controllers

import (
	"net/http"
	"application/models"
)

type HomeController struct {
	name string
}

func (h *HomeController) Index(res http.ResponseWriter, req *http.Request) {
	user := models.User{"Bjørn-Inge Morstad", "bimorstad@gmail.com"}
	renderView(defaultLayout, "home", "index", user, res, req)
}

