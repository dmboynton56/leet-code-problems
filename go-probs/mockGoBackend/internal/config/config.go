package config

import "os"

// Config holds runtime settings. A small struct like this is typical in Go backends:
// static types make misconfiguration obvious at compile time rather than at runtime.
type Config struct {
	Port string
}

func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return Config{Port: port}
}
