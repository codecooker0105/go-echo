package main

import (
	"github.com/joho/godotenv"
	"github.com/triaton/forum-backend-echo/server"
	"log"
)

func main() {
	// Load environment file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s := server.MakeServer()
	s.Logger.Fatal(s.Start(":1200"))
}
