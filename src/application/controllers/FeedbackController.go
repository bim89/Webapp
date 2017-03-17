package controllers

import (
	"net/http"
	"application/models"
	"encoding/json"
	"fmt"
)

type FeedbackController struct {
}

func (*FeedbackController) Create(res http.ResponseWriter, req *http.Request) {
	f := models.Feedback{}
	json.NewDecoder(req.Body).Decode(&f)
	msg, err := f.Save()
	if err != nil {
		fmt.Fprint(res, "Error:%s", err.Error())
	} else {
		fmt.Fprintf(res, msg)
	}

}