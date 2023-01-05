package configwarp

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
type ConfigSensors struct {
	Sensor []configSensor
}

func (cfSensors *ConfigSensors) ReadConfig(location string) {
	data, err := os.ReadFile(location)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &cfSensors.Sensor)
	if err != nil {
		panic(err)
	}
}

type ConfigDB struct {
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

func (db *ConfigDB) ReadConfig(location string) {
	data, err := os.ReadFile(location)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &db)
	if err != nil {
		panic(err)
	}
}

type ConfigLine struct {
	Token   string  `json:"token"`
	Message lineMsg `json:"message"`
}
type lineMsg struct {
	Alarm      string `json:"alarm"`
	BackNormal string `json:"back_normal"`
}

func (cfLine *ConfigLine) ReadConfig(location string) {
	data, err := os.ReadFile(location)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &cfLine)
	if err != nil {
		panic(err)
	}
}

type ConfigBuzzer struct {
	Name     string `json:"name"`
	Port     string `json:"port"`
	Baud     int    `json:"baud"`
	ModbusID int    `json:"modbus_id"`
}

func (cfBuzzer *ConfigBuzzer) ReadConfig(location string) {
	data, err := os.ReadFile(location)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &cfBuzzer)
	if err != nil {
		panic(err)
	}
}
