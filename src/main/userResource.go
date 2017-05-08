package main

import (
	"github.com/gorilla/mux"
	"application/controllers"
)

func userResource(r *mux.Router) {
	u := controllers.UserController{}
	r.HandleFunc("/create", u.Create).Methods("POST")
	r.HandleFunc("/login", u.Login).Methods("POST")
	r.HandleFunc("/logout", u.Logout).Methods("POST")
}