package controllers

import (
	"net/http"
	"encoding/json"
	"application/models"
	"fmt"
	"github.com/golang/glog"
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


func (*UserTestController) Show(res http.ResponseWriter, req *http.Request) {
	id := req.FormValue("t")
	ut := models.UserTest{}
	results := ut.FindId(id)

	json.NewEncoder(res).Encode(results)
}


func (*UserTestController) Delete(res http.ResponseWriter, req *http.Request) {
	ut := models.UserTest{}
	message, err := ut.Delete(req.FormValue("t"))
	fmt.Fprintf(res, message)
	if err != nil {
		glog.Warning(err)
	}
}