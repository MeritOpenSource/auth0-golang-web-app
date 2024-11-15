package main

import (
	"01-Login/platform/merit"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"01-Login/platform/authenticator"
	"01-Login/platform/router"
)

func main() {
	if err := godotenv.Load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}
	meritClient, err := merit.New()
	if err != nil {
		log.Fatalf("Failed to initialize the merit client: %v", err)
	}

	rtr := router.New(auth, meritClient)

	log.Print("Server listening on http://localhost:3000/")
	if err := http.ListenAndServe("0.0.0.0:3000", rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
