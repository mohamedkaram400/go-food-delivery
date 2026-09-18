package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port              string
	IdentityService   string
	RestaurantService string
	OrderService      string
	DeliveryService   string
	NotificationService   string
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		Port:              		os.Getenv("API_GATEWAY_SERVICE"),
		IdentityService:   		os.Getenv("IDENTITY_SERVICE"),
		RestaurantService: 		os.Getenv("RESTAURANT_SERVICE"),
		OrderService:      		os.Getenv("OORDER_SERVICE"),
		DeliveryService:   	   	os.Getenv("DELIVERY_SERVICE"),
		NotificationService:   	os.Getenv("NOTIFICATION_SERVICE"),
	}
}