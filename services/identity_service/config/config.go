package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)


type Config struct {
	Port              string
	IdentityService   string
	DSN   			  string
	MigrationPath	  string
	DatabaseURL		  string
	JWTSecret		  string
	AccessTokenDuration		time.Duration
	RefreshTokenDuration	time.Duration
}

func Load() *Config {
	_ = godotenv.Load()

	accessDuration, err := time.ParseDuration(os.Getenv("ACCESS_TOKEN_DURATION"))
    if err != nil {
        log.Fatal("Invalid ACCESS_TOKEN_DURATION: ", err)
    }

    refreshDuration, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_DURATION"))
    if err != nil {
        log.Fatal("Invalid REFRESH_TOKEN_DURATION: ", err)
    }

	jwtSecret := os.Getenv("SECRET_KEY")
    if jwtSecret == "" {
        log.Fatal("SECRET_KEY is missing")
    }


	log.Println("accessDuration: ", accessDuration, "refreshDuration: ", refreshDuration)
	return &Config{
		Port:              getOrDefault(os.Getenv("SERVICE_PORT"), ":50051"),
		IdentityService:   os.Getenv("IDENTITY_SERVICE_HOST"),
		DatabaseURL: 	   os.Getenv("DATABASE_URL"),
		MigrationPath:     os.Getenv("MIGRATION_PATH"),
		DSN:   			   os.Getenv("DSN"),
		JWTSecret:   	   os.Getenv("SECRET_KEY"),
		AccessTokenDuration:      accessDuration,
		RefreshTokenDuration:     refreshDuration,
	}
}

func getOrDefault(key string, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return def
}

