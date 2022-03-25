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
	SearchRadius  int
}

func NewService(driverAPIAddr string, searchRadius int) (*Service, error) {
	return &Service{
		DriverAPIAddr: driverAPIAddr,
		SearchRadius:  searchRadius,
	}, nil
}

func (s *Service) FindDriver(rider model.GeoPoint) (*model.DriverInfo, error) {
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

func (s *Service) findSuitableDrivers(rider model.GeoPoint) ([]*model.DriverInfo, error) {

	// Anonymous for request body
	req := struct {
		Rider  model.GeoPoint `json:"rider"`
		Radius int            `json:"radius"`
	}{
		Rider:  rider,
		Radius: s.SearchRadius,
	}
	rider_data, err := json.Marshal(req)

	if err != nil {
		logrus.WithError(err).Error("error marshalling rider data")
		return nil, err
	}

	// make post request to driver API
	resp, err := http.Post(s.DriverAPIAddr, "application/json",
		bytes.NewBuffer(rider_data))

	if err != nil {
		logrus.WithError(err).Error("error getting drivers info")
		return nil, err
	}
	defer resp.Body.Close()

	// parse response
	var drivers []*model.DriverInfo
	err = json.NewDecoder(resp.Body).Decode(&drivers)
	if err != nil {
		logrus.WithError(err).Error("error getting drivers info")
		return nil, err
	}
	// return suitable drivers
	return drivers, nil

}

func searchMostSuitable(drivers []*model.DriverInfo) *model.DriverInfo {

	var mostSuitableDriver *model.DriverInfo
	var minDistance float64
	for _, driver := range drivers {

		err := driver.Validate()
		if err != nil {
			continue
		}
		distance := driver.Distance
		if mostSuitableDriver == nil || distance < minDistance {
			minDistance = distance
			mostSuitableDriver = driver
		}
	}

	return mostSuitableDriver
}
