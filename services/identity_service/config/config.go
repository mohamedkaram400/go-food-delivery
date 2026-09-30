package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)


type Config struct {
	Port              string
	IdentityService   string
	DSN   			  string
	MigrationPath	  string
	DatabaseURL		  string
	AccessTokenDuration		int
	RefreshTokenDuration	int
}

func Load() *Config {
	_ = godotenv.Load()

	accessTokenDuration, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_DURATION"))
	refreshTokenDuration, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_DURATION"))

	return &Config{
		Port:              getOrDefault(os.Getenv("SERVICE_PORT"), ":50051"),
		IdentityService:   os.Getenv("IDENTITY_SERVICE_HOST"),
		DatabaseURL: 	   os.Getenv("DATABASE_URL"),
		MigrationPath:     os.Getenv("MIGRATION_PATH"),
		DSN:   			   os.Getenv("DSN"),
		AccessTokenDuration:      accessTokenDuration,
		RefreshTokenDuration:     refreshTokenDuration,
	}
}

func getOrDefault(key string, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return def
}

