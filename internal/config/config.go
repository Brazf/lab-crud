package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser  string
	DBPass  string
	DBHost  string
	DBPort  string
	DBName  string
	AppPort string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("Não foi possível carregar as variáveis de ambiente!")
	}

	return &Config{
		DBUser:  os.Getenv("DB_USER"),
		DBPass:  os.Getenv("DB_PASS"),
		DBHost:  os.Getenv("DB_HOST"),
		DBPort:  os.Getenv("DB_PORT"),
		DBName:  os.Getenv("DB_NAME"),
		AppPort: os.Getenv("APP_PORT"),
	}
}
