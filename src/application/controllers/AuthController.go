package controllers

import (
	"net/http"
	"github.com/satori/go.uuid"
	"log"
	"application/config"
	"application/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/haisum/recaptcha"
	"fmt"
)

type AuthController struct {
}

func (*AuthController) Create(res http.ResponseWriter, req *http.Request) {
	re := recaptcha.R{
		Secret: "6LcjAyQUAAAAAB3fdbvUbiUPYNoG5dgeJrWebzre",
	}
	email := req.FormValue("email")
	password := req.FormValue("password")
	log.Print(email)
	log.Print(password)

	isValid := re.Verify(*req)
	if isValid {

		var a= models.Admin{}
		a = a.GetByEmail(email)

		if err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password)); err == nil {
			session, err := config.Sessions().Get(req, "AUTH")
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			uid := uuid.NewV4()
			session.Values["email"] = a.Email
			session.Values["uuid"] = uid.String()
			session.Values["loggedIn"] = true

			if err := session.Save(req, res); err != nil {
				log.Panic(err.Error())
			} else {
				http.Redirect(res, req, "/", 301)
			}
		} else {
			http.Redirect(res, req, "/login", 301)
		}
	} else {
		fmt.Fprintf(res, "Invalid! These errors ocurred: %v", re.LastError())
	}
}

func (*AuthController) Delete(res http.ResponseWriter, req *http.Request) {
	session, err := config.Sessions().Get(req, "AUTH")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["loggedIn"] = false
	session.Values["email"] = ""
	session.Values["uuid"] = ""
	if err := session.Save(req, res); err != nil {
		log.Panic(err.Error())
	} else {
		http.Redirect(res, req, "/login", 301)
	}
}
