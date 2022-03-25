package match

import (
	"reflect"
	"testing"
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
