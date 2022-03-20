package model

import (
	"reflect"
	"testing"
)

func TestGeoPoint_Validate(t *testing.T) {
	type fields struct {
		Lat  float64
		Long float64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "valid GeoPoint", fields: fields{Lat: -180, Long: -90}, wantErr: false}, {name: "valid GeoPoint", fields: fields{Lat: 180, Long: -90}, wantErr: false}, {name: "valid GeoPoint", fields: fields{Lat: 22, Long: -22}, wantErr: false}, {name: "invalid GeoPoint", fields: fields{Lat: -181, Long: 91}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GeoPoint{
				Lat:  tt.fields.Lat,
				Long: tt.fields.Long,
			}
			if err := g.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("GeoPoint.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewGeoPoint(t *testing.T) {
	type args struct {
		lat  float64
		long float64
	}
	tests := []struct {
		name string
		args args
		want *GeoPoint
	}{
		{name: "valid GeoPoint", args: args{lat: -180, long: -90}, want: &GeoPoint{Lat: -180, Long: -90}}, {name: "valid GeoPoint", args: args{lat: 180, long: -90}, want: &GeoPoint{Lat: 180, Long: -90}}, {name: "valid GeoPoint", args: args{lat: 22, long: -22}, want: &GeoPoint{Lat: 22, Long: -22}}, {name: "invalid GeoPoint", args: args{lat: -181, long: 91}, want: nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGeoPoint(tt.args.lat, tt.args.long); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGeoPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
