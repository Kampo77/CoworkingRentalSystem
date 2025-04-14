package config

type Config struct {
	OrderServiceURL     string
	InventoryServiceURL string
}

func LoadConfig() *Config {
	return &Config{
		OrderServiceURL:     "http://localhost:8082",
		InventoryServiceURL: "http://localhost:8081",
	}
}
