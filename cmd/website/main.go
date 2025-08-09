package main

import (
	"log"

	"github.com/gsvd/website/internal/app"
	"github.com/joho/godotenv"
)

// Set at build time via ldflags
var CommitHash = "dev"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env: %v", err)
	}

	app, err := app.New(CommitHash)
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	app.RegisterRoutes()

	if err := app.Start(); err != nil {
		log.Fatalf("server exited with error: %v", err)
	}
}
