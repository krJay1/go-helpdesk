package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type dbConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

type Config struct {
	*dbConfig
	JWTSecret string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env found")
	}

	return &Config{
		dbConfig: &dbConfig{
			DBHost:     os.Getenv("DB_HOST"),
			DBPort:     os.Getenv("DB_PORT"),
			DBUser:     os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_NAME"),
		},
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}
