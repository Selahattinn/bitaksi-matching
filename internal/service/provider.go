package service

import (
	"github.com/Selahattinn/bitaksi-matching/internal/service/match"
)

type Provider struct {
	cfg          *Config
	matcgService *match.Service
}

func NewProvider(cfg *Config) (*Provider, error) {
	matchService, err := match.NewService(cfg.DriverAPIAddr, cfg.SearchRadius)
	if err != nil {
		return nil, err
	}
	return &Provider{
		cfg:          cfg,
		matcgService: matchService,
	}, nil
}

func (p *Provider) GetConfig() *Config {
	return p.cfg
}

func (p *Provider) GetMatchService() *match.Service {
	return p.matcgService
}
