package main

import (
	"log"

	"github.com/Gsvd/website/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env: %v", err)
	}

	app, err := app.New()
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	app.RegisterRoutes()

	if err := app.Start(); err != nil {
		log.Fatalf("server exited with error: %v", err)
	}
}
