package main

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	AllowedOrigin   string
	InstagramSecret string
}

func NewConfig() *Config {
	godotenv.Load()

	cfg := Config{
		Port:            os.Getenv("PORT"),
		AllowedOrigin:   os.Getenv("ALLOWED_ORIGIN"),
		InstagramSecret: os.Getenv("INSTAGRAM_CLIENT_SECRET"),
	}

	return &cfg
}
