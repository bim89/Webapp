package main

import (
	"application/controllers"
	"github.com/gorilla/mux"
)

func userTestResource(r *mux.Router) {
	ut := controllers.UserTestController{}
	r.HandleFunc("/create", ut.Create).Methods("POST")
	r.HandleFunc("/read", ut.Read).Methods("GET")
	r.HandleFunc("/delete", ut.Delete).Methods("GET")
	// r.HandleFunc("/innstillinger", Home.Settings).Methods("GET")
}