package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Selahattinn/bitaksi-matching/internal/api/response"
	"github.com/Selahattinn/bitaksi-matching/internal/model"
)

// FindDriver godoc
// @Summary      Finds a driver
// @Description  Get driver by GeoPoint
// @Tags         GeoPoints
// @Accept       json
// @Produce      json
// @Param        GeoPoint   body      model.GeoPoint  true  "An GeoPoint object with json format"
// @Success      200  {object}  model.GeoPoint
// @Failure      404  {object}  response.Error
// @Router       /findDriver [post]
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
