package config

import "os"

var (
	Port           string
	Environment    string
	LogLevel       string
	AllowedOrigins string
)

func Load() {
	var ok bool

	Port, ok = os.LookupEnv("PORT")
	if !ok || Port == "" {
		Port = "8080"
	}

	Environment, ok = os.LookupEnv("ENVIRONMENT")
	if !ok || Environment == "" {
		Environment = "production"
	}

	LogLevel, ok = os.LookupEnv("LOG_LEVEL")
	if !ok || LogLevel == "" {
		LogLevel = "info"
	}

	AllowedOrigins, ok = os.LookupEnv("ALLOWED_ORIGINS")
	if !ok || AllowedOrigins == "" {
		panic("ALLOWED_ORIGINS is not set")
	}
}
