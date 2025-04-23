package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB        DBConfig
	Redis     string
	JWTSecret string
}

type DBConfig struct {
	Name     string
	User     string
	Host     string
	Password string
	Port     string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to load .env %w", err)
	}
	cfg := &Config{

		DB: DBConfig{
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Host:     os.Getenv("DB_HOST"),
			Password: os.Getenv("DB_PASSWORD"),
			Port:     os.Getenv("DB_PORT"),
		},
		Redis:     os.Getenv("REDIS_URL"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	if cfg.DB.Name == "" || cfg.DB.User == "" || cfg.DB.Host == "" || cfg.DB.Port == "" {
		return nil, fmt.Errorf("missing required DB environmental variables")
	}
	if cfg.Redis == "" {
		return nil, fmt.Errorf("missing Redis_URL")
	}
	return cfg, nil
}
