package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	DBName        string
	DBUserName    string
	DBPassword    string
	DBHost        string
	DBAddress     string
	JWTExpiration int64
	JWTSecret     string
}

var EnvVars = initializeConfig()

func initializeConfig() Config {

	godotenv.Load()

	return Config{
		DBName:        GetEnvAsString("DB_NAME", "goDB"),
		DBUserName:    GetEnvAsString("DB_USER_NAME", "root"),
		DBPassword:    GetEnvAsString("DB_PASSWORD", "root"),
		DBHost:        GetEnvAsString("PUBLIC_HOST", "http://localhost"),
		DBAddress:     fmt.Sprintf("%s:%s", GetEnvAsString("DB_HOST", "127.0.01"), GetEnvAsString("DB_PORT", "3306")),
		Port:          GetEnvAsString("PORT", "8080"),
		JWTExpiration: GetEnvAsInt("JWT_EXP", 3600*24*7),
		JWTSecret:     GetEnvAsString("JWT_SECRET", "top-secret-key"),
	}
}

func GetEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
func GetEnvAsString(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
