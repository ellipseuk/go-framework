package core

type Config struct {
	Port string
}

// LoadConfig loads the application configuration
func LoadConfig() *Config {
	return &Config{
		Port: "8080",
	}
}
