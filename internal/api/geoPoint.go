package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Selahattinn/bitaksi-matching/internal/api/response"
	"github.com/Selahattinn/bitaksi-matching/internal/model"
)

// Match godoc
// @Summary      Finds a driver
// @Description  Get driver by GeoPoint
// @Tags         Match
// @Accept       json
// @Produce      json
// @Param        GeoPoint   body      model.GeoPoint  true  "An GeoPoint object with json format"
// @Success      200  {object}  model.GeoPoint "Driver's geoPoint"
// @Failure      404  {object}  response.Error "Not found"
// @Router       /match [post]
func (a *API) Match(w http.ResponseWriter, r *http.Request) {
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
