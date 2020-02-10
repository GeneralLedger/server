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

	db, err := database.NewConnection(
		os.Getenv("POSTGRES_URL"),
		database.Config{
			MigrationsDir: "./database/migrations",
		},
	)
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
