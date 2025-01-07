package config

import "os"

type Config struct {
	DBName     string
	DBUserName string
	DBPassword string
	DBHost     string
}

var EnvVars = initializeConfig()

func initializeConfig() Config {
	return Config{
		DBName:     GetEnv("DB_NAME", "test"),
		DBUserName: GetEnv("DB_USER_NAME", "root"),
		DBPassword: GetEnv("DB_PASSWORD", "root"),
		DBHost:     GetEnv("DB_HOST", "localhost"),
	}
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
