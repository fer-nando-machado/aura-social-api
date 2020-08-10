package main

import (
	"errors"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port string `envconfig:"AURASOCIAL_PORT" default:"8888"`
}

func NewConfig() (*Config, error) {
	cfg := Config{}

	err := envconfig.Process("AURASOCIAL", &cfg)
	if err != nil {
		return nil, errors.New("error to process env variables")
	}

	return &cfg, nil
}
