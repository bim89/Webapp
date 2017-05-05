package main

import (
	"application/controllers"
	"github.com/gorilla/mux"
	"net/http"
	"application/config"
)


type Val struct {
	loggedIn bool
}


func loggedIn(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := config.Sessions().Get(r, "AUTH")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if session.Values["loggedIn"] == true {
			next(w, r)
		} else {
			redirect(w, r)
		}
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
	r.HandleFunc("/login", Home.Login).Methods("GET")
}