package main

import (
	// "application/models"
	"application/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	var homeController = controllers.HomeController{}
	router.HandleFunc("/", homeController.Index)


	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println("Listening on localhost:8080...")
	/*
	fmt.Println("Hello World")
	another()

	var userController = controllers.UserController{}
	userController.Create()
	controllers.Post.Create()

	arr := [5]int{1, 2, 3, 4, 5}
	for _, x := range arr {
		fmt.Printf("The value is: %d and som more text \n", x)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i + 1)
	}

	models.Print()
	*/
}
