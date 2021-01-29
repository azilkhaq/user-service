package main

import (
	"log"
	"os"
	"user-service/controllers"

	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	server.Run(os.Getenv("SERVER_PORT"))
}