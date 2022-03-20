package api

import (
	"net/http"

	"github.com/Selahattinn/bitaksi-matching/internal/api/response"
	"github.com/gorilla/mux"
)

// API configuration
type Config struct {
	SigningSecret string `yaml:"signing_secret"`
}

// API represents the structure of the API
type API struct {
	Router *mux.Router
	config *Config
}

// New returns the api settings
func New(cfg *Config, router *mux.Router) (*API, error) {
	api := &API{
		config: cfg,
		Router: router,
	}

	// Endpoint for browser preflight requests
	api.Router.Methods("OPTIONS").HandlerFunc(api.corsMiddleware(api.preflightHandler))

	// Endpoint for healtcheck
	api.Router.HandleFunc("/api/v1/health", api.corsMiddleware(api.logMiddleware(api.healthHandler))).Methods("GET")

	// Endpoint for geoPoints
	api.Router.HandleFunc("/api/v1/findDriver", api.corsMiddleware(api.logMiddleware(api.authMiddleware(api.FindDriver)))).Methods("POST")

	return api, nil
}

func (a *API) healthHandler(w http.ResponseWriter, r *http.Request) {
	response.Write(w, r, struct {
		Status string `json:"status"`
	}{
		"ok",
	})

}

func (a *API) preflightHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

}
