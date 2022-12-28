package jsonReader

import (
	"reflect"
	"testing"
)

func Test_jsonMap_Read(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		j       *jsonMap
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "json string",
			j:       &jsonMap{},
			args:    args{input: `[{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"},{"name": "battery sensor", "capacity": 50, "time": "2019-01-21T19:07:28Z"}]`},
			wantErr: false,
		},
		{
			name:    "json string",
			j:       &jsonMap{},
			args:    args{input: `[{"name": battery sensor, "capacity": 40, "time": "2019-01-21T19:07:28Z"},{"name": "battery sensor", "capacity": 50, "time": "2019-01-21T19:07:28Z"}]`},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.j.Read(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("jsonMap.Read() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewjsonMap(t *testing.T) {
	tests := []struct {
		name string
		want *jsonMap
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewjsonMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonMap_ReadFile(t *testing.T) {
	type args struct {
		location string
	}
	tests := []struct {
		name    string
		j       *jsonMap
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			j:    &jsonMap{},
			args: args{
				location: "../config/sensor.json",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.j.ReadFile(tt.args.location); (err != nil) != tt.wantErr {
				t.Errorf("jsonMap.ReadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
