package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	routes := mux.NewRouter().StrictSlash(true)
	dbHandler()

	homeResource(routes)
	userTestResource(routes.PathPrefix("/usertest/").Subrouter())

	// Including assets:
	routes.PathPrefix("/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./src/application/assets/"))))
	fmt.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", routes))
}
