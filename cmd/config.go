package main

import (
	"fmt"
	// 	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string // You can add more fields as needed
}

func loadConfig() (*Config, error) {
	cfg, err := godotenv.Read(".env") // Optional: loads .env file automatically
	if err != nil {
		return nil, fmt.Errorf("failed to read .env file: %v", err)
	}

	// 	dbURL := os.Getenv("DATABASE_URL")
	// 	if dbURL == "" {
	// 		return nil, fmt.Errorf("DATABASE_URL is not set")
	// 	}

	return &Config{
		DatabaseURL: cfg["DATABASE_URL"],
		Port:        cfg["PORT"],
	}, nil
}
