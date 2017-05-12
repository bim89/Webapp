package controllers

import (
	"net/http"
	"application/models"
	"fmt"
	"github.com/satori/go.uuid"
	"encoding/json"
	"strconv"
)

type UserController struct {
}


func (*UserController) Create(res http.ResponseWriter, req *http.Request) {
	u := models.User{}
	u.Email = req.FormValue("email")
	u.Username = req.FormValue("username")
	age,err := strconv.ParseInt(req.FormValue("age"), 10, 8)
	if (err == nil) {
		u.Age = int(age)
	}
	u.Gender = req.FormValue("gender")
	if (req.FormValue("password") == req.FormValue("confirmPassword")) {
		u.Password = req.FormValue("password")
		u.UUID = uuid.NewV4().String()
		u.Save()
		json.NewEncoder(res).Encode(u)
	} else {
		fmt.Fprint(res, "Passwords did not match")
	}
}

func (*UserController) ConfirmUUID(res http.ResponseWriter, req *http.Request) {
	u := models.User{}
	u.Email = req.FormValue("email")
	u,hasMail := u.CheckEmail(u.Email)
	if (hasMail && u.UUID == req.FormValue("uuid")) {
		fmt.Fprint("Confirmed")
	} else {
		fmt.Fprint("Wrong data given")
	}
}

func (*UserController) Login(res http.ResponseWriter, req *http.Request) {
	email := req.FormValue("email")
	password := req.FormValue("password")
	var u = models.User{}
	u, hasMail := u.CheckEmail(email)
	if (hasMail) {
		if (u.Email == email && u.Password == password) {
			u.UUID = uuid.NewV4().String()
			u.Update(u)
			msg := fmt.Sprintf("{'UUID': '%s', 'msg': 'You have been logged in'}", u.UUID)
			fmt.Fprint(res, msg)
		} else {
			fmt.Fprint(res, "Wrong password")
		}
	} else {
		fmt.Fprint(res, "Email not registered")
	}
}


func (*UserController) Logout(res http.ResponseWriter, req *http.Request) {
	u := models.User{}
	email := req.FormValue("email")
	uuid := req.FormValue("uuid")
	u, hasMail := u.CheckEmail(email)
	if (hasMail) {
		if (u.Email == email && u.UUID == uuid) {
			u.UUID = ""
			u.Update(u)
			fmt.Fprint(res, "You have been logged out")
		} else {
			fmt.Fprint(res, "UUID did not match with given email")
		}
	} else {
		fmt.Fprint(res, "The given email does not exist")
	}
}




