package main

import (
	"testing"
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
				wSensors: &configSensors{},
				wDB:      &configDB{},
				wLine:    &configLine{},
				wBuzzer:  &configBuzzer{},
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

func Test_work_LineTrigger(t *testing.T) {
	type args struct {
		sendMsg string
	}
	tests := []struct {
		name string
		w    *work
		args args
	}{
		// TODO: Add test cases.
		{
			name: "line notify send",
			w: &work{
				wSensors: &configSensors{},
				wDB:      &configDB{},
				wLine: &configLine{
					Token: "",
					Message: lineMsg{
						Alarm:      "",
						BackNormal: "",
					},
				},
				wBuzzer: &configBuzzer{},
			},
			args: args{
				sendMsg: "this is test for golang yongfeng temperature version",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.LineTrigger(tt.args.sendMsg)
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
				wSensors: &configSensors{},
				wDB:      &configDB{},
				wLine:    &configLine{},
				wBuzzer: &configBuzzer{
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
