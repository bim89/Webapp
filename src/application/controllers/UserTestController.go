package controllers

import (
	"net/http"
	"encoding/json"
	"application/models"
)

type UserTestController struct {
}

func (*UserTestController) Create(res http.ResponseWriter, req *http.Request) {
	ut := models.UserTest{}
	json.NewDecoder(req.Body).Decode(&ut)
	ut.Save()
}

func (*UserTestController) Read(res http.ResponseWriter, req *http.Request) {
	ut := models.UserTest{}
	list := ut.FindAll()

	// res.Header().Set("Content-Type", "application/vnd.api+sjon")
	json.NewEncoder(res).Encode(list)


}
