package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	
	Host       string
	Port       string
	DbUsername string
	DbPassword string
	DbName     string
	DbSslMode  string
	
}

func GetConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Fetch values from environment
	config := &Config{
		Port:       os.Getenv("PORT"),
		Host:       os.Getenv("HOST"),
		DbUsername: os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		DbSslMode:  os.Getenv("DB_SSLMODE"),
	}

	if config.Port == "" || config.Host == "" || config.DbUsername == "" ||
		config.DbPassword == "" || config.DbName == "" || config.DbSslMode == "" {
		log.Fatal("Missing required environment variables. Check your .env file.")
	}

	log.Println("Config loaded successfully!")
	return config
}
