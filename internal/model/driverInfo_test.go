package model

import (
	"reflect"
	"testing"
)

func TestNewDriverInfo(t *testing.T) {
	type args struct {
		geoPoint GeoPoint
		distance float64
	}
	tests := []struct {
		name string
		args args
		want *DriverInfo
	}{
		{name: "valid DriverInfo", args: args{geoPoint: GeoPoint{Lat: -180, Long: -90}, distance: 0}, want: &DriverInfo{GeoPoint: GeoPoint{Lat: -180, Long: -90}, Distance: 0}}, {name: "valid DriverInfo", args: args{geoPoint: GeoPoint{Lat: 180, Long: -90}, distance: 0}, want: &DriverInfo{GeoPoint: GeoPoint{Lat: 180, Long: -90}, Distance: 0}}, {name: "valid DriverInfo", args: args{geoPoint: GeoPoint{Lat: 22, Long: -22}, distance: 0}, want: &DriverInfo{GeoPoint: GeoPoint{Lat: 22, Long: -22}, Distance: 0}}, {name: "invalid DriverInfo", args: args{geoPoint: GeoPoint{Lat: -181, Long: 91}, distance: 0}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDriverInfo(tt.args.geoPoint, tt.args.distance); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDriverInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriverInfo_Validate(t *testing.T) {
	type fields struct {
		GeoPoint GeoPoint
		Distance float64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "valid DriverInfo", fields: fields{GeoPoint: GeoPoint{Lat: -180, Long: -90}, Distance: 0}, wantErr: false}, {name: "valid DriverInfo", fields: fields{GeoPoint: GeoPoint{Lat: 180, Long: -90}, Distance: 0}, wantErr: false}, {name: "valid DriverInfo", fields: fields{GeoPoint: GeoPoint{Lat: 22, Long: -22}, Distance: 0}, wantErr: false}, {name: "invalid DriverInfo", fields: fields{GeoPoint: GeoPoint{Lat: -181, Long: 91}, Distance: 0}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DriverInfo{
				GeoPoint: tt.fields.GeoPoint,
				Distance: tt.fields.Distance,
			}
			if err := d.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("DriverInfo.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
