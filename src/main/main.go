package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"application/controllers"
)

func main() {

	routes := mux.NewRouter().StrictSlash(true)
	homeResource(routes)

	userTestResource(routes.PathPrefix("/usertest/").Subrouter())
	a := controllers.AuthController{}
	routes.HandleFunc("/authenticate", a.Create).Methods("POST")
	// routes.HandleFunc("/logout",).Methods("POST")

	feedbackResource(routes.PathPrefix("/feedback/").Subrouter())

	// Including assets:
	routes.PathPrefix("/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./src/application/assets/"))))
	fmt.Println("Listening on localhost:8081")
	log.Fatal(http.ListenAndServe(":8001", routes))
}
