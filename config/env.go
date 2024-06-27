package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBName     string
	DBPassword string
	DBHost     string
	DBPort     string
}

var Envs = initConfig()

func initConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Something went wrong while loading env", err)
	}
	return Config{
		DBUser:     getEnv("DB_USER", "postgres"),
		DBName:     getEnv("DB_NAME", "geoecommerce"),
		DBPassword: getEnv("DB_PASSWORD", "milano"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
	}

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
