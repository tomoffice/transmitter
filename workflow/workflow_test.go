package workflow

import (
	"testing"
	PkgConfig "transmitter/configwrap"
)

func Test_work_GetConfig(t *testing.T) {
	tests := []struct {
		name string
		w    *work
	}{
		// TODO: Add test cases.
		{
			name: "",
			w: &work{
				wSensors: &PkgConfig.ConfigSensors{},
				wDB:      &PkgConfig.ConfigDB{},
				wLine:    &PkgConfig.ConfigLine{},
				wBuzzer:  &PkgConfig.ConfigBuzzer{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.GetConfig()
		})
	}
}

func Test_work_GetValue(t *testing.T) {
	tests := []struct {
		name string
		w    *work
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.GetValue()
		})
	}
}

func Test_work_BuzzerTrigger(t *testing.T) {
	tests := []struct {
		name string
		w    *work
	}{
		// TODO: Add test cases.
		{
			name: "",
			w: &work{
				wSensors: &PkgConfig.ConfigSensors{},
				wDB:      &PkgConfig.ConfigDB{},
				wLine:    &PkgConfig.ConfigLine{},
				wBuzzer: &PkgConfig.ConfigBuzzer{
					Name:     "simulate relay",
					Port:     "COM3",
					Baud:     9600,
					ModbusID: 100,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.BuzzerTrigger()
		})
	}
}
