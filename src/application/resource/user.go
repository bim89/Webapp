package resource

import (
	"application/controllers"
	"github.com/gorilla/mux"
)

type UserResource struct {
}

func (resource *UserResource) routes() {
	resource.route.HandleFunc("/", controllers.Home.Index).Methods("GET")

}
