package match

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Selahattinn/bitaksi-matching/internal/model"
	"github.com/sirupsen/logrus"
)

var (
	ErrDriverNotFound = errors.New("driver not found")
)

type Service struct {
	DriverAPIAddr string
	MaxDistance   float64
	DriverAPIKey  string
}

func NewService(driverAPIAddr string, maxDistance float64, driverAPIKey string) (*Service, error) {
	return &Service{
		DriverAPIAddr: driverAPIAddr,
		MaxDistance:   maxDistance,
		DriverAPIKey:  driverAPIKey,
	}, nil
}

func (s *Service) FindDriver(rider model.GeoPoint) (*model.SearchResult, error) {
	suitableDrivers, err := s.findSuitableDrivers(rider)
	if err != nil {
		return nil, err
	}
	if suitableDrivers != nil {
		driver := searchMostSuitable(suitableDrivers)
		if driver != nil {
			return driver, err
		}
		return nil, ErrDriverNotFound
	}
	return nil, ErrDriverNotFound
}

func (s *Service) findSuitableDrivers(rider model.GeoPoint) ([]*model.SearchResult, error) {

	// Anonymous for request body
	reqData := struct {
		Rider  model.GeoPoint `json:"rider"`
		Radius float64        `json:"max_distance"`
	}{
		Rider:  rider,
		Radius: s.MaxDistance,
	}
	rider_data, err := json.Marshal(reqData)

	if err != nil {
		logrus.WithError(err).Error("error marshalling rider data")
		return nil, err
	}

	// make post request to driver API
	request, err := http.NewRequest("POST", s.DriverAPIAddr, bytes.NewBuffer(rider_data))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", s.DriverAPIKey)
	if err != nil {
		logrus.WithError(err).Error("error getting drivers info")
		return nil, err
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		logrus.WithError(err).Error("error getting drivers info")
		return nil, err
	}
	defer resp.Body.Close()

	// parse response
	var drivers []*model.SearchResult
	err = json.NewDecoder(resp.Body).Decode(&drivers)
	if err != nil {
		logrus.WithError(err).Error("error getting drivers info")
		return nil, err
	}

	// return suitable drivers
	return drivers, nil

}

func searchMostSuitable(drivers []*model.SearchResult) *model.SearchResult {

	var mostSuitableDriver *model.SearchResult
	var minDistance float64
	for _, driver := range drivers {
		distance := driver.Distance
		if mostSuitableDriver == nil || distance < minDistance {
			minDistance = distance
			mostSuitableDriver = driver
		}
	}

	return mostSuitableDriver
}
