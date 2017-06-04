package controllers

import (
	"net/http"
	"application/models"
	"application/config"
	"fmt"
	"encoding/json"

)

type HomeController struct {
}

func (*HomeController) Index(res http.ResponseWriter, req *http.Request) {
	// renderView(res, req, nil)
	ut := models.UserTest{}
	session, err := config.Sessions().Get(req, "AUTH")
	if (err != nil) {
		fmt.Fprint(res, err.Error())
	}
	email := session.Values["email"].(string)
	usertests := ut.FindAll(email, true)

	data, jerr := json.Marshal(usertests)
	if jerr != nil {

	}

	createTemplate("index", "home", "layout", string(data), res, req)
}

func (HomeController) Create(res http.ResponseWriter, req *http.Request) {
	// renderView(res, req, nil)
	session, err := config.Sessions().Get(req, "AUTH")
	if (err != nil) {
		fmt.Fprint(res, err.Error())
	} else {
		createTemplate("create", "home", "layout", session.Values, res, req)
	}
}

func (HomeController) Show(res http.ResponseWriter, req *http.Request) {
	id := req.FormValue("t")
	createTemplate("show", "home", "layout", id, res, req)
}

func (*HomeController) Settings(res http.ResponseWriter, req *http.Request) {
	// renderView(res, req, nil)
	createTemplate("settings", "home", "layout", nil, res, req)
}

func (*HomeController) Login(res http.ResponseWriter, req *http.Request) {
	// renderView(res, req, nil)
	// createTemplate("login", "home", "", nil, res, req)
	sitekey := "6LcjAyQUAAAAAP3yvsBhQxbNb25FfbwmCsIDmAy7"

	createTemplate("login_form", "home", "login", sitekey, res, req)
}
