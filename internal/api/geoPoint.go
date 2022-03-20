package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Selahattinn/bitaksi-matching/internal/api/response"
	"github.com/Selahattinn/bitaksi-matching/internal/model"
)

func (a *API) FindDriver(w http.ResponseWriter, r *http.Request) {
	var geoPoint model.GeoPoint
	err := json.NewDecoder(r.Body).Decode(&geoPoint)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting GeoPoint info: %v", err), http.StatusNotFound, "")
		return
	}

	// Valide geoPoint
	err = geoPoint.Validate()
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting Validating geoPoint: %v", err), http.StatusNotFound, "")
		return
	}

}
