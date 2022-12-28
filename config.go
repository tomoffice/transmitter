package main

import (
	"encoding/json"
	"os"
)

type configSensor struct {
	SensorID    int         `json:"id"`
	Name        string      `json:"name"`
	StationID   int         `json:"station_id"`
	Port        string      `json:"port"`
	Baud        int         `json:"baud"`
	SensorType  string      `json:"sensor_type"`
	DataType    string      `json:"data_type"`
	ModbusType  string      `json:"modbus_type"`
	ModbusCode  string      `json:"modbus_code"`
	DbID        int         `json:"db_id"`
	ValueScale  string      `json:"value_scale"`
	Alert       []float64   `json:"alert"`
	AlertStatus bool        `json:"-"`
	Value       interface{} `json:"-"`
}
type configSensors struct {
	Sensor []configSensor
}

func (cfSensors *configSensors) ReadConfig(location string) {
	data, err := os.ReadFile(location)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &cfSensors.Sensor)
	if err != nil {
		panic(err)
	}
}

type configDB struct {
	Type       string `json:"type"`
	DBusername string `json:"dbusername"`
	DBpassword string `json:"dbpassword"`
	Protocol   string `json:"protocol"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Database   string `json:"database"`
	Charset    string `json:"charset"`
	Location   string `json:"location"`
	Parsetime  bool   `json:"parsetime"`
}

func (db *configDB) ReadConfig(location string) {
	data, err := os.ReadFile(location)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &db)
	if err != nil {
		panic(err)
	}
}

type configLine struct {
	Token   string  `json:"token"`
	Message lineMsg `json:"message"`
}
type lineMsg struct {
	Alarm      string `json:"alarm"`
	BackNormal string `json:"back_normal"`
}

func (cfLine *configLine) ReadConfig(location string) {
	data, err := os.ReadFile(location)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &cfLine)
	if err != nil {
		panic(err)
	}
}

type configBuzzer struct {
	Name     string `json:"name"`
	Port     string `json:"port"`
	Baud     int    `json:"baud"`
	ModbusID int    `json:"modbus_id"`
}

func (cfBuzzer *configBuzzer) ReadConfig(location string) {
	data, err := os.ReadFile(location)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &cfBuzzer)
	if err != nil {
		panic(err)
	}
}
