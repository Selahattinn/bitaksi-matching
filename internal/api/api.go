// @title           Matching API
// @version         1.0
// @description     This is a matching server.
// @termsOfService  https://www.selahattinceylan.com

// @contact.name   API Support
// @contact.url    https://www.selahattinceylan.com
// @contact.email  selahattinceylan9622@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
// @SecurityDefinitions.ApiKey Bearer
// @in header
// @name Authorization
// @type apiKey
package api

import (
	"net/http"

	_ "github.com/Selahattinn/bitaksi-matching/docs"
	"github.com/Selahattinn/bitaksi-matching/internal/api/response"
	"github.com/Selahattinn/bitaksi-matching/internal/service"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// API configuration
type Config struct {
	SigningSecret string `yaml:"signing_secret"`
}

// API represents the structure of the API
type API struct {
	Router  *mux.Router
	config  *Config
	Service service.Service
}

// New returns the api settings
func New(cfg *Config, router *mux.Router, svc service.Service) (*API, error) {
	api := &API{
		config:  cfg,
		Router:  router,
		Service: svc,
	}

	// Endpoint for browser preflight requests
	api.Router.Methods("OPTIONS").HandlerFunc(api.corsMiddleware(api.preflightHandler))

	// Endpoint for healtcheck
	api.Router.HandleFunc("/api/v1/health", api.corsMiddleware(api.logMiddleware(api.healthHandler))).Methods("GET")

	api.Router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

	// Endpoint for geoPoints
	api.Router.HandleFunc("/api/v1/match", api.corsMiddleware(api.logMiddleware(api.authMiddleware(api.Match)))).Methods("POST")

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
