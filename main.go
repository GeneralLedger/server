package main

import (
	"database/sql"
	"github.com/generalledger/api/server"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"os"
)

func main() {
	godotenv.Load()
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
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
