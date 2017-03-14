package main

import (
	"github.com/gorilla/mux"
	"application/controllers"
)

func answersResource(r *mux.Router) {
	a := controllers.AnswersController{}
	r.HandleFunc("/create", a.Create).Methods("POST")
}
