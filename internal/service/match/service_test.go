package match

import (
	"reflect"
	"testing"

	"github.com/Selahattinn/bitaksi-matching/internal/model"
)

func TestNewService(t *testing.T) {
	type args struct {
		driverAPIAddr string
		maxDistance   float64
		driverAPIKey  string
	}
	tests := []struct {
		name    string
		args    args
		want    *Service
		wantErr bool
	}{
		{name: "success creation", args: args{driverAPIAddr: "http://localhost:8081", maxDistance: 100}, want: &Service{DriverAPIAddr: "http://localhost:8081", MaxDistance: 100}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewService(tt.args.driverAPIAddr, tt.args.maxDistance, tt.args.driverAPIKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchMostSuitable(t *testing.T) {
	type args struct {
		drivers []*model.SearchResult
	}
	tests := []struct {
		name string
		args args
		want *model.SearchResult
	}{
		{name: "success search", args: args{drivers: []*model.SearchResult{&model.SearchResult{Distance: 1, Driver: model.Driver{ID: 1, Location: model.GeoJson{Type: "Point", Coordinates: []float64{10, 10}}}}, &model.SearchResult{Distance: 14, Driver: model.Driver{ID: 1, Location: model.GeoJson{Type: "Point", Coordinates: []float64{10, 10}}}}}}, want: &model.SearchResult{Distance: 1, Driver: model.Driver{ID: 1, Location: model.GeoJson{Type: "Point", Coordinates: []float64{10, 10}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMostSuitable(tt.args.drivers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchMostSuitable() = %v, want %v", got, tt.want)
			}
		})
	}
}
