package controllers

import (
	"net/http"
	"github.com/satori/go.uuid"
	"log"
	"application/config"
)

type AuthController struct {
}

func (*AuthController) Create(res http.ResponseWriter, req *http.Request) {

	session, err := config.Sessions().Get(req, "AUTH")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	uid := uuid.NewV4()
	session.Values["uuid"] = uid.String()
	println(uid.String())
	println("AUTH")
	session.Values["loggedIn"] = true

	if err := session.Save(req, res); err != nil {
		log.Panic(err.Error())
	}


}

func (*AuthController) Delete(res http.ResponseWriter, req *http.Request) {

}
