package controllers

import (
	"net/http"
	"application/models"
	"encoding/json"
)

type FeedbackController struct {
}

func (*FeedbackController) Create(res http.ResponseWriter, req *http.Request) {
	f := models.Feedback{}
	json.NewDecoder(req.Body).Decode(&f)
	f.Save()
}