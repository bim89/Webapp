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
			next(w, r)
		} else {
			redirect(w, r)
		}

		log.Println("Executing middlewareTwo again")
	})
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", 301)
}

func homeResource(r *mux.Router) {
	Home := controllers.HomeController{}

	r.HandleFunc("/", loggedIn(Home.Index)).Methods("GET")
	r.HandleFunc("/brukerundersokelse", loggedIn(Home.Show)).Methods("GET")
	r.HandleFunc("/ny-undersokelse", loggedIn(Home.Create)).Methods("GET")
	r.HandleFunc("/innstillinger", loggedIn(Home.Settings)).Methods("GET")
	r.HandleFunc("/login", Home.Login).Methods("GET")
}