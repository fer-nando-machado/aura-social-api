package main

import (
		"os"
    "github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func NewConfig() *Config {
	godotenv.Load()

	cfg := Config{
		Port: os.Getenv("PORT"),
	}

	return &cfg
}
