package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetVariable unexported
func GetVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// Config unexported
type Config struct {
	Email    string
	Password string
	Port     string
}

// Init exported
func Init() Config {
	Email := GetVariable("EMAIL")
	Password := GetVariable("EMAIL_PASSWORD")
	Port := GetVariable("PORT")
	return Config{Email, Password, Port}
}
