package routes

import (
	"github.com/gildemberg-santos/crud-mvc-golang/controller"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.Use(controller.LoggingMiddleware)

	r.HandleFunc("/user", controller.UserCreateController).Methods("POST")
	r.HandleFunc("/users", controller.UserFindAllController).Methods("GET")
	r.HandleFunc("/user/{id}", controller.UserUpdateController).Methods("PUT")
	r.HandleFunc("/user/{id}", controller.UserFindByIdController).Methods("GET")
	r.HandleFunc("/user/{id}", controller.UserDeleteByIdController).Methods("DELETE")

	return r
}
