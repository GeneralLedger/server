package main

import (
	"github.com/generalledger/api/database"
	"github.com/generalledger/api/server"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"os"
)

func main() {
	godotenv.Load()

	db, err := database.MigratedDatabase(os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	server := &server.Server{
		Port:   os.Getenv("PORT"),
		Router: httprouter.New(),
		DB:     db,
	}
	server.Start()
}
