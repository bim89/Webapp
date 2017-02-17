package main

import (
	"application/controllers"
	"github.com/gorilla/mux"
)

func userResource(r *mux.Router) {
	Home := controllers.HomeController{}
	r.HandleFunc("/", Home.Index).Methods("GET")
	r.HandleFunc("/create", Home.Create).Methods("POST")
}
