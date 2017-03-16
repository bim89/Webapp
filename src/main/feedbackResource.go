package main

import (
	"github.com/gorilla/mux"
	"application/controllers"
)

func feedbackResource(r *mux.Router) {
	a := controllers.FeedbackController{}
	r.HandleFunc("/create", a.Create).Methods("POST")
}
