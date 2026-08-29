package config

type Config struct {
	Port              string
	IdentityService   string
	DSN   			  string
}

func Load() *Config {
	// if os.Getenv("APP_ENV") != "production" {
	// 	_ = godotenv.Load()
	// }

	return &Config{
		Port:              ":50051",
		IdentityService:   "localhost:50051",
		DSN:   			   "root:qazwsx123@tcp(localhost:3306)/identity_service?charset=utf8mb4&parseTime=True&loc=Local",
	}
}