package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

type Config struct {
	Email    string
	Password string
	Port     string
}

func Init() Config {
	Email := GetVariable("EMAIL")
	Password := GetVariable("EMAIL_PASSWORD")
	Port := GetVariable("PORT")
	return Config{Email, Password, Port}
}
