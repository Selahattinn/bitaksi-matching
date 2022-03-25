package server

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/Selahattinn/bitaksi-matching/internal/api"
	"github.com/Selahattinn/bitaksi-matching/internal/service"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ListenAddress string `yaml:"host"`
	// For HTTPS
	//CertFile      string `yaml:"cert_file"`
	//KeyFile       string `yaml:"key_file"`
	API     *api.Config     `yaml:"api"`
	Service *service.Config `yaml:"service"`
}

// Instance represents an instance of the server
type Instance struct {
	Config     *Config
	API        *api.API
	httpServer *http.Server
	Service    service.Service
}

// NewInstance returns an new instance of our server
func NewInstance(cfg *Config) (*Instance, error) {
	var router = mux.NewRouter()

	// Initialize ServiceProvider
	serviceProvider, err := service.NewProvider(cfg.Service)
	if err != nil {
		logrus.WithError(err).Fatal("Could not initialize Service")
		return nil, err
	}
	// Initialize API
	api, err := api.New(cfg.API, router, serviceProvider)
	if err != nil {
		logrus.WithError(err).Fatal("Could not initialize API")
		return nil, err
	}

	return &Instance{
		Config:  cfg,
		API:     api,
		Service: serviceProvider,
	}, nil
}

// Start starts the server
func (i *Instance) Start() {

	// Startup the HTTP Server in a way that we can gracefully shut it down again
	i.httpServer = &http.Server{
		Addr:    i.Config.ListenAddress,
		Handler: i.API.Router,
	}

	err := i.httpServer.ListenAndServe()
	if err != http.ErrServerClosed {
		logrus.WithError(err).Error("HTTP Server stopped unexpected")
		i.Shutdown()
	} else {
		logrus.WithError(err).Info("HTTP Server stopped")
	}
}

// Shutdown stops the server
func (i *Instance) Shutdown() {
	// Shutdown HTTP server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := i.httpServer.Shutdown(ctx)
	if err != nil {
		logrus.WithError(err).Error("Failed to shutdown HTTP server gracefully")
		os.Exit(1)
	}

	logrus.Info("Shutdown HTTP server...")
	os.Exit(0)
}
