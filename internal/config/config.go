package config

import (
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"os"
)

var config *Config

// App is config for app
type App struct {
	Port string
	Env  string
}

// Cache is config for cache
type Cache struct {
	Host string
	Port string
}

// Config is config struct
type Config struct {
	App   App
	Cache Cache
}

// GetConfig returns app config
func GetConfig() (*Config, error) {
	if config == nil {
		err := godotenv.Load()
		if err != nil {
			return nil, errors.Wrap(err, "could not load app-config")
		}

		config = &Config{
			App: App{
				Port: os.Getenv("APP_PORT"),
				Env:  os.Getenv("APP_ENV"),
			},
			Cache: Cache{
				Host: os.Getenv("REDIS_HOST"),
				Port: os.Getenv("REDIS_PORT"),
			},
		}
	}

	return config, nil
}
