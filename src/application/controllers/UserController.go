package controllers

import (
	"net/http"
	"application/models"
	"fmt"
)

type UserController struct {
}


func (*UserController) Create(res http.ResponseWriter, req *http.Request) {

}

func (*UserController) Login(res http.ResponseWriter, req *http.Request) {
	email := req.FormValue("email")
	password := req.FormValue("password")
	var u = models.User{}
	u, hasMail := u.CheckEmail(email)
	if (hasMail) {
		if (u.Email == email && u.Password == password) {
			fmt.Fprint(res, "You have been logged in")
		} else {
			fmt.Fprint(res, "Wrong password")
		}
	} else {
		fmt.Fprint(res, "Email not registered")
	}
}


func (*UserController) Logout(res http.ResponseWriter, req *http.Request) {

}




