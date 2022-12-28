package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"text/template"
	"time"
	PkgEval "transmitter/evalwapper"
	PkgModbus "transmitter/modbus"
	PkgSql "transmitter/sql"
	Pkgtools "transmitter/tools"
)

type work struct {
	wSensors *configSensors
	wDB      *configDB
	wLine    *configLine
	wBuzzer  *configBuzzer
}

// Workflower for interface
type Workflower interface {
	GetConfig()
	GetValue()
	LineTrigger(string)
	BuzzerTrigger()
	Notify()
	WriteDB()
	NotifyReport()
}

func (w *work) GetConfig() {
	fmt.Println("!GetConfig!")
	sensorsConf := &configSensors{}
	sensorsConf.ReadConfig("config/sensor.json")
	w.wSensors = sensorsConf

	dbConf := &configDB{}
	dbConf.ReadConfig("config/database.json")
	w.wDB = dbConf

	lineConf := &configLine{}
	lineConf.ReadConfig("config/line.json")
	w.wLine = lineConf

	buzzerConf := &configBuzzer{}
	buzzerConf.ReadConfig("config/buzzer.json")
	w.wBuzzer = buzzerConf
}
func (w *work) GetValue() {
	fmt.Println("-------------", time.Now(), "-------------")
	fmt.Println("!GetValue!")
	for i, v := range w.wSensors.Sensor {
		switch v.ModbusType {
		case "rtu":
			var rtu PkgModbus.Modbus = PkgModbus.NewRtuCmd()
			rtu.SetCmd(func(cmd string) []byte {
				data := strings.ReplaceAll(cmd, " ", "")
				output, err := hex.DecodeString(data)
				if err != nil {
					fmt.Println(err)
				}
				return output
			}(v.ModbusCode))
			var Serial PkgModbus.Transmit = PkgModbus.NewRSerial(v.Name, v.Port, v.Baud, time.Millisecond*1000)
			Serial.Conn()
			Serial.Flush()
			Serial.Write(rtu.GetCmd())
			data, err := Serial.Read()
			if Pkgtools.NewHandle().Err(err, func() {
				log.Printf("%s -> %v\n", v.Name, err)
			}, time.Second*0) {
				Serial.Disconn()
				continue
			}
			Serial.Disconn()
			value, err := rtu.Unpack(data, strToType(v.DataType))
			if Pkgtools.NewHandle().Err(err, func() {}, time.Second*0) {
				continue
			}
			transData, err := PkgEval.Eval(map[string]interface{}{
				"value": value,
			}, v.ValueScale)

			if Pkgtools.NewHandle().Err(err, func() {}, time.Second*0) {
				continue
			}
			if data == nil {
				w.wSensors.Sensor[i].Value = nil
			} else {
				w.wSensors.Sensor[i].Value = transData.(float64)
			}

			log.Printf("%s -> %.2f\n", v.Name, transData)
			time.Sleep(time.Millisecond * 500)
		case "tcp":
			//sensor := PkgModbus.NewNetWork()
		}
	}
}
func (w *work) Notify() {
	fmt.Println("!Notify!")
	var buzzerOnce bool = false
	for i, v := range w.wSensors.Sensor {
		var min float64 = v.Alert[0]
		var max float64 = v.Alert[1]
		var now string = time.Now().Format("2006-01-02 15:04:05")
		var alertStatus bool = w.wSensors.Sensor[i].AlertStatus //copy
		if v.Value != nil {
			var value float64 = v.Value.(float64)
			if value < 100 && value > -10 {
				if value <= min || value >= max {
					if alertStatus == false {
						vars := map[string]interface{}{
							"value":       v.Value,
							"time":        now,
							"sensor_name": v.Name,
						}
						var buf bytes.Buffer
						t := template.Must(template.New("").Parse(w.wLine.Message.Alarm))
						err := t.Execute(&buf, vars)
						if err != nil {
							panic(err)
						}
						if Pkgtools.NewHandle().Err(err, func() {}, time.Second*0) {
							continue
						}
						w.LineTrigger(buf.String())
						if buzzerOnce == false {
							w.BuzzerTrigger()
							buzzerOnce = true
						}
						w.wSensors.Sensor[i].AlertStatus = true //ptr
					}
				} else {
					if alertStatus == true {
						vars := map[string]interface{}{
							"value":       v.Value,
							"time":        now,
							"sensor_name": v.Name,
						}
						var buf bytes.Buffer
						t := template.Must(template.New("").Parse(w.wLine.Message.BackNormal))
						err := t.Execute(&buf, vars)
						if err != nil {
							panic(err)
						}
						w.LineTrigger(buf.String())
						w.wSensors.Sensor[i].AlertStatus = false //ptr
					}
				}

			} else {
				w.LineTrigger("溫度計讀取異常請查修,不開啟警報器!!!" + fmt.Sprintf("%.2f", v.Value))
			}
		}
	}
}
func (w *work) WriteDB() {
	fmt.Println("!WriteDB!")
	switch w.wDB.Type {
	case "mysql":
		var DB PkgSql.Database = PkgSql.NewMysql()
		DB.ParseConfig(map[string]interface{}{
			"dbusername": w.wDB.DBusername,
			"dbpassword": w.wDB.DBpassword,
			"protocol":   w.wDB.Protocol,
			"host":       w.wDB.Host,
			"port":       w.wDB.Port,
			"database":   w.wDB.Database,
			"charset":    w.wDB.Charset,
			"location":   w.wDB.Location,
			"parsetime":  w.wDB.Parsetime,
		})
		err := DB.Conn()
		if Pkgtools.NewHandle().Retry(err, 5, time.Second*1, DB.Conn, func() {
			fmt.Println("database alive")
		}, func() {
			fmt.Println("database dead already")
		}) {
			return
		}
		for _, v := range w.wSensors.Sensor {
			if v.Value != nil {
				err := DB.Insert(fmt.Sprintf("INSERT INTO record VALUES (NULL,%d,%.2f,NOW())", v.DbID, v.Value))
				if err != nil {
					continue
				}
			}
		}

		defer DB.Close()
		log.Println("Write DB success")
	}
}
func (w *work) LineTrigger(sendMsg string) {
	fmt.Println("!LineTrigger!")
	url := fmt.Sprintf("https://notify-api.line.me/api/notify?message=%s", url.QueryEscape(sendMsg))
	method := "POST"
	payload := strings.NewReader("")
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", w.wLine.Token))
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(body))
	log.Printf("Line Notify Send %s", sendMsg)
}
func (w *work) BuzzerTrigger() {
	fmt.Println("!BuzzerTrigger!")
	var serial PkgModbus.Transmit = PkgModbus.NewRSerial(
		w.wBuzzer.Name,
		w.wBuzzer.Port,
		w.wBuzzer.Baud,
		time.Millisecond*200)
	var Cmd PkgModbus.Modbus = PkgModbus.NewRtuCmd()
	Cmd.SetCmd([]byte{byte(w.wBuzzer.ModbusID), 0x05, 0x00, 0x00, 0xFF, 0x00})
	onCmd := Cmd.GetCmd()
	Cmd.SetCmd([]byte{byte(w.wBuzzer.ModbusID), 0x05, 0x00, 0x00, 0x00, 0x00})
	offCmd := Cmd.GetCmd()
	serial.Conn()
	serial.Write(onCmd)
	time.Sleep(time.Second)
	serial.Write(offCmd)
	serial.Flush()
	serial.Disconn()
}
func (w *work) NotifyReport() {

	var reportTime string = time.Now().Format("2006-01-02 15:04:05")
	var reportTime1 string = time.Now().Format("01-02 15:04")
	var reportStr string

	reportStr += fmt.Sprintf("\n%18s", "每日更新報")
	reportStr += fmt.Sprintf("\n%-3s%15s%3s", "***", reportTime, "***")
	reportStr += fmt.Sprintf("\n%s%4s%9s", "冰庫", "溫度", "最後監測")
	for _, v := range w.wSensors.Sensor {
		if v.Value == nil {
			reportStr += fmt.Sprintf("\n%3d%7s%14s", v.DbID, "無資料", reportTime1)
		} else {
			reportStr += fmt.Sprintf("\n%3d%8.1f%16s", v.DbID, v.Value, reportTime1)
		}

	}
	w.LineTrigger(reportStr)
}
func strToType(strType string) reflect.Kind {
	switch strType {
	case "int8": //1byte
		return reflect.Int8
	case "int16": //2byte
		return reflect.Int16
	case "int32": //4byte
		return reflect.Int32
	case "int64": //8byte
		return reflect.Int64
	case "uint8": //1byte
		return reflect.Uint8
	case "uint16": //2byte
		return reflect.Uint16
	case "uint32": //4byte
		return reflect.Uint32
	case "uint64": //8byte
		return reflect.Uint64
	case "float32": //2byte
		return reflect.Float32
	case "float64": //4byte
		return reflect.Float64
	default:
		panic("unknow type")
	}
}
