package model

type SearchResult struct {
	Driver   Driver  `json:"driver"`
	Distance float64 `json:"distance"` // in km
}

type Driver struct {
	ID       int64   `bson:"_id" json:"id"`
	Location GeoJson `bson:"location" json:"location"`
}

// swagger:model
type GeoJson struct {
	Type string `json:"type"`

	// example: [-122.083739,37.423021]
	Coordinates []float64 `json:"coordinates"`
}
