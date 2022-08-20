package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvStripeApiKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("STRIPE_API_KEY")
}
