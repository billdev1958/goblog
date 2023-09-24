package main

import (
	"log"
	"os"

	"github.com/billdev1958/goblog.git/api"
	"github.com/billdev1958/goblog.git/db"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	store, err := db.NewPostgreStore()
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(os.Getenv("PORT"), store)
	server.Run()
}
