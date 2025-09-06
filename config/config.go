package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	SenderEmail   string
	Port          string
	RedisAddr     string
	RedisPassword string
	RedisDB       int
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	// Load .env file if it exists
	_ = godotenv.Load()

	config := &Config{
		SenderEmail:   os.Getenv("SENDER_EMAIL"),
		Port:          os.Getenv("PORT"),
		RedisAddr:     os.Getenv("REDIS_ADDR"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
	}

	// Parse Redis DB
	if redisDBStr := os.Getenv("REDIS_DB"); redisDBStr != "" {
		if db, err := strconv.Atoi(redisDBStr); err == nil {
			config.RedisDB = db
		}
	}

	// Set default values
	if config.Port == "" {
		config.Port = "3000"
	}
	if config.RedisAddr == "" {
		config.RedisAddr = "localhost:6379"
	}

	// Validate required environment variables
	if config.SenderEmail == "" {
		log.Fatal("You must set the SENDER_EMAIL environment variable to your verified sender email address.")
	}

	return config
}
