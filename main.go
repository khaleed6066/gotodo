package main

import (
	"crudApp/config"
	"crudApp/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Working on crud")
	config.InitDB()
	defer config.DBConnect.Close()
	// Create router
	router := mux.NewRouter()

	routes.RegisterTodoRoutes(router)

	// Starting the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
