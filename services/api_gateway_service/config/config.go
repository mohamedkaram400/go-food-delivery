package config

type Config struct {
	Port              string
	IdentityService   string
	RestaurantService string
	OrderService      string
	DeliveryService   string
	NotificationService   string
}

func Load() *Config {
	return &Config{
		Port:              ":8080",
		IdentityService:   "localhost:50051",
		RestaurantService: "localhost:50052",
		OrderService:      "localhost:50053",
		DeliveryService:   	   "localhost:50054",
		NotificationService:   "localhost:50055",
	}
}