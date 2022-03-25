package service

import "github.com/Selahattinn/bitaksi-matching/internal/service/match"

type Config struct {
	SearchRadius  int    `yaml:"search_radius"`
	DriverAPIAddr string `yaml:"driver_api_addr"`
}

type Service interface {
	GetConfig() *Config
	GetMatchService() *match.Service
}
