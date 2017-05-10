package main

import (
	"github.com/gorilla/mux"
	"application/controllers"
)

func authResource(r *mux.Router) {
	a := controllers.AuthController{}
	r.HandleFunc("/create", a.Create).Methods("POST")
	r.HandleFunc("/delete", a.Delete).Methods("POST")
}
