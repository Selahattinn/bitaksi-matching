package service

import (
	"github.com/Selahattinn/bitaksi-matching/internal/service/match"
)

type Provider struct {
	cfg          *Config
	matchService *match.Service
}

func NewProvider(cfg *Config) (*Provider, error) {
	matchService, err := match.NewService(cfg.DriverAPIAddr, cfg.MaxDistance, cfg.DriverAPIKey)
	if err != nil {
		return nil, err
	}
	return &Provider{
		cfg:          cfg,
		matchService: matchService,
	}, nil
}

func (p *Provider) GetConfig() *Config {
	return p.cfg
}

func (p *Provider) GetMatchService() *match.Service {
	return p.matchService
}
