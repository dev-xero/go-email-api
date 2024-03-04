package main

import (
	"dev-xero/email-api/application"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[ERROR]: Unable to load environment variables")
	}

	mux := http.NewServeMux()
	app := application.Application{}

	app.New(mux)
	app.Initialize()

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	log.Println("[INFO]: Server is running on http://localhost:8080")
	server.ListenAndServe()
}
