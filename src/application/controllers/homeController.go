package controllers

import (
	"application/models"
	"net/http"
	"fmt"
)

type HomeController struct {
}

func (*HomeController) Index(res http.ResponseWriter, req *http.Request) {
	user := models.User{}
	user.Name = "Bjørn-Inge Morstad"
	user.Email = "bimorstad@gmail.com"

	renderView(res, req, user)
}

func (*HomeController) Create(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Post method")
}
