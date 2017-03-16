package main

import (
	"application/controllers"
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

func loggedIn(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c,_ := r.Cookie("loggedIn")
		if c != nil {
			redirect(w, r)
		} else {
			next(w, r)
		}

		log.Println("Executing middlewareTwo again")
	})
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", 301)
}

func homeResource(r *mux.Router) {
	Home := controllers.HomeController{}

	r.HandleFunc("/", Home.Index).Methods("GET")
	r.HandleFunc("/brukerundersokelse", Home.Show).Methods("GET")
	r.HandleFunc("/ny-undersokelse", Home.Create).Methods("GET")
	r.HandleFunc("/innstillinger", Home.Settings).Methods("GET")
}