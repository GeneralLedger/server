package main

import (
	"github.com/generalledger/api/server"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"os"
)

func main() {
	godotenv.Load()
	server := &server.Server{
		Port:   os.Getenv("PORT"),
		Router: httprouter.New(),
	}
	server.Start()
}
