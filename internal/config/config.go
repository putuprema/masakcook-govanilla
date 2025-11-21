// Package config for server configurations
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	Environment string
	StaticDir   string
}

func Load() *Config {
	_ = godotenv.Load()
	return &Config{
		Port:        getEnv("PORT", "8000"),
		Environment: getEnv("ENVIRONMENT", "development"),
		StaticDir:   getEnv("STATIC_DIR", "./static"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func (c *Config) IsDevelopment() bool {
	return c.Environment == "development"
}

func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}

func (c *Config) LogConfig() {
	log.Printf("Configuration loaded:")
	log.Printf("  Environment: %s", c.Environment)
	log.Printf("  Port: %s", c.Port)
	log.Printf("  Static Dir: %s", c.StaticDir)
}
