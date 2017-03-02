package main

import (
	"application/controllers"
	"github.com/gorilla/mux"
)

func homeResource(r *mux.Router) {
	Home := controllers.HomeController{}
	r.HandleFunc("/", Home.Index).Methods("GET")
	r.HandleFunc("/ny-undersokelse", Home.Create).Methods("GET")
	r.HandleFunc("/innstillinger", Home.Settings).Methods("GET")
}