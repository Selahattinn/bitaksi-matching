package service

import "github.com/Selahattinn/bitaksi-matching/internal/service/match"

type Config struct {
	MaxDistance   float64 `yaml:"max_distance"`
	DriverAPIAddr string  `yaml:"driver_api_addr"`
	DriverAPIKey  string  `yaml:"driver_api_key"`
}

type Service interface {
	GetConfig() *Config
	GetMatchService() *match.Service
}
