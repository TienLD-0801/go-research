package main

import (
	configs "go-backend/internal/configs"

	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := configs.NewServer()
	server.Start()
}
