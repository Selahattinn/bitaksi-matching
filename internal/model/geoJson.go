package model

import (
	"errors"
)

var (
	ErrInvalidLatitude  = errors.New("invalid latitude")
	ErrInvalidLongitude = errors.New("invalid longitude")
)

type GeoPoint struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

func NewGeoPoint(lat float64, long float64) *GeoPoint {
	if lat < -180 || lat > 180 {
		return nil
	}
	if long < -90 || long > 90 {
		return nil
	}
	return &GeoPoint{
		Lat:  lat,
		Long: long,
	}
}

func (g *GeoPoint) Validate() error {
	if g.Lat < -180 || g.Lat > 180 {
		return ErrInvalidLatitude
	}
	if g.Long < -90 || g.Long > 90 {
		return ErrInvalidLongitude
	}
	return nil
}
