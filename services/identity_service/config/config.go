package config

type Config struct {
	Port              string
	IdentityService   string
}

func Load() *Config {
	// if os.Getenv("APP_ENV") != "production" {
	// 	_ = godotenv.Load()
	// }

	return &Config{
		Port:              ":50051",
		IdentityService:   "localhost:50051",
	}
}