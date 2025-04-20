package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort      string
	DBHost       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBPort       string
	Env          string
	JWTSecret    string
	ClientKey    string
	ClientSecret string
}

var AppConfig Config

func LoadConfig() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("no .env file found")
	}

	AppConfig = Config{
		AppPort:      getEnv("APP_PORT", "3000"),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBPassword:   getEnv("DB_PASSWORD", "password"),
		DBName:       getEnv("DB_NAME", "postgres"),
		DBPort:       getEnv("DB_PORT", "5432"),
		Env:          getEnv("GO_ENV", "development"),
		JWTSecret:    getEnv("JWT_SECRET", "secret"),
		ClientKey:    getEnv("CLIENT_KEY", "key"),
		ClientSecret: getEnv("CLIENT_SECRET", "secret"),
	}

	return nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
