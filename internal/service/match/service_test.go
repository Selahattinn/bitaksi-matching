package match

import (
	"reflect"
	"testing"

	"github.com/Selahattinn/bitaksi-matching/internal/model"
)

func TestNewService(t *testing.T) {
	type args struct {
		driverAPIAddr string
		searchRadius  int
	}
	tests := []struct {
		name    string
		args    args
		want    *Service
		wantErr bool
	}{
		{name: "success creation", args: args{driverAPIAddr: "http://localhost:8081", searchRadius: 100}, want: &Service{DriverAPIAddr: "http://localhost:8081", SearchRadius: 100}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewService(tt.args.driverAPIAddr, tt.args.searchRadius)
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
		drivers []*model.DriverInfo
	}
	tests := []struct {
		name string
		args args
		want *model.DriverInfo
	}{
		{name: "success search", args: args{drivers: []*model.DriverInfo{&model.DriverInfo{GeoPoint: model.GeoPoint{Lat: 12, Long: 45}, Distance: 12}, &model.DriverInfo{GeoPoint: model.GeoPoint{Lat: 12, Long: 45}, Distance: 12}, &model.DriverInfo{GeoPoint: model.GeoPoint{Lat: 12, Long: 45}, Distance: 22}}}, want: &model.DriverInfo{GeoPoint: model.GeoPoint{Lat: 12, Long: 45}, Distance: 12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMostSuitable(tt.args.drivers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchMostSuitable() = %v, want %v", got, tt.want)
			}
		})
	}
}
