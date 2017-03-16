package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	routes := mux.NewRouter().StrictSlash(true)
	homeResource(routes)

	ut := routes.PathPrefix("/usertest/").Subrouter()
	userTestResource(ut)
	f := routes.PathPrefix("/feedback/").Subrouter()
	feedbackResource(f)

	// Including assets:
	routes.PathPrefix("/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./src/application/assets/"))))
	fmt.Println("Listening on localhost:8081")
	log.Fatal(http.ListenAndServe("localhost:8001", routes))
}
