package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBName     string
	DBUserName string
	DBPassword string
	DBHost     string
	DBAddress  string
}

var EnvVars = initializeConfig()

func initializeConfig() Config {

	godotenv.Load()

	return Config{
		DBName:     GetEnv("DB_NAME", "goDB"),
		DBUserName: GetEnv("DB_USER_NAME", "root"),
		DBPassword: GetEnv("DB_PASSWORD", "root"),
		DBHost:     GetEnv("PUBLIC_HOST", "http://localhost"),
		DBAddress:  fmt.Sprintf("%s:%s", GetEnv("DB_HOST", "127.0.01"), GetEnv("DB_PORT", "3306")),
		Port:       GetEnv("PORT", "8080"),
	}
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
