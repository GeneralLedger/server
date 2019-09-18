package main

import (
	"database/sql"
	"fmt"
	"github.com/generalledger/api/server"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"os"
)

func main() {
	godotenv.Load()
	_, err := Connection(os.Getenv("POSTGRES_URL"))
	if err != nil {
		fmt.Println(err)
	}
	server := &server.Server{
		Port:   os.Getenv("PORT"),
		Router: httprouter.New(),
	}
	server.Start()
}

// Connection opens a new database connection
// to the specified database URL.
func Connection(databaseURL string) (*sql.DB, error) {
	// Open the DB connection
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	// Ensure we can ping the database
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// OK
	return db, nil
}
