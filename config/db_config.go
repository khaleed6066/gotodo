package config

import (
	"log"

	"github.com/gocraft/dbr/v2"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

var DBConnect *dbr.Connection

func InitDB() {
	var err error

	// Data Source Name (DSN) for PostgreSQL
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=todo_db sslmode=disable"

	// Initialize the DB connection
	DBConnect, err = dbr.Open("postgres", dsn, nil)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	log.Println("Connected to PostgreSQL successfully!")
}
