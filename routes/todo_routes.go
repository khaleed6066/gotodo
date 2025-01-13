package routes

import (
	"crudApp/handlers"

	"github.com/gorilla/mux"
)

func RegisterTodoRoutes(router *mux.Router) {

	router.HandleFunc("/todos", handlers.AllTodoHandler).Methods("GET")
	router.HandleFunc("/todos/{id}", handlers.GetTodoHandler).Methods("GET")
	router.HandleFunc("/todos", handlers.CreateTodoHandler).Methods("POST")
	router.HandleFunc("/todos/{id}", handlers.UpdateTodoHandler).Methods("PUT")
	router.HandleFunc("/todos/{id}", handlers.DeleteTodoHandler).Methods("DELETE")

}
