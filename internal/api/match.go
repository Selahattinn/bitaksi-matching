package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Selahattinn/bitaksi-matching/internal/api/response"
	"github.com/Selahattinn/bitaksi-matching/internal/model"
)

// Match godoc
// @Summary      Match a driver
// @Description  Get rider by GeoPoint and find the most suitable driver
// @Tags         Match
// @Accept       json
// @Produce      json
// @Param        GeoPoint   body      model.GeoPoint  true  "An GeoPoint object with json format"
// @Success      200  {object}  model.SearchResult "Result of the search"
// @Failure      404  "Not found"
// @Failure      401  "UnAuthorized"
// @Router       /match [post]
// @Security Bearer
func (a *API) Match(w http.ResponseWriter, r *http.Request) {
	var riderGeoPoint model.GeoPoint
	err := json.NewDecoder(r.Body).Decode(&riderGeoPoint)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting GeoPoint info: %v", err), http.StatusNotFound, "")
		return
	}

	// Valide geoPoint
	err = riderGeoPoint.Validate()
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting Validating geoPoint: %v", err), http.StatusNotFound, "")
		return
	}

	//Find Rider
	driver, err := a.Service.GetMatchService().FindDriver(riderGeoPoint)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting Driver: %v", err), http.StatusNotFound, "")
		return
	}
	response.Write(w, r, driver)

}
