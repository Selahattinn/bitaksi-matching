package model

import "errors"

var (
	ErrInvalidDistance = errors.New("invalid distance")
)

type DriverInfo struct {
	GeoPoint GeoPoint `json:"geoPoint"`
	Distance float64  `json:"distance"`
}

func NewDriverInfo(geoPoint GeoPoint, distance float64) *DriverInfo {
	err := geoPoint.Validate()
	if err != nil {
		return nil
	}

	if distance < 0 {
		return nil
	}
	return &DriverInfo{
		GeoPoint: geoPoint,
		Distance: distance,
	}
}

func (d *DriverInfo) Validate() error {
	if d.Distance < 0 {
		return ErrInvalidDistance
	}
	err := d.GeoPoint.Validate()
	if err != nil {
		return err
	}
	return nil
}
