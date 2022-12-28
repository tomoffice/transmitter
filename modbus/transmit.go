package modbus

import (
	"errors"
	"net"
	"time"

	"github.com/tarm/serial"
)

// Transmit interface
type Transmit interface {
	Conn()
	Write([]byte)
	Read() ([]byte, error)
	Flush() error
	Disconn()
}

// Serial for RS485...etc
type cserial struct {
	name    string
	config  *serial.Config
	sserial *serial.Port
}

func NewRSerial(name string, port string, baud int, readtimeout time.Duration) *cserial {

	return &cserial{
		name: name,
		config: &serial.Config{
			Name:        port,
			Baud:        baud,
			ReadTimeout: readtimeout,
			Size:        serial.DefaultSize,
			Parity:      serial.ParityNone,
			StopBits:    serial.Stop1,
		},
	}
}
func (RS *cserial) Conn() {

	serial, err := serial.OpenPort(RS.config)
	if err != nil {
		panic(err)
	}
	RS.sserial = serial
}
func (RS *cserial) Write(b []byte) {
	RS.sserial.Write(b)
}
func (RS *cserial) Read() ([]byte, error) {
	buf := make([]byte, 1024)
	len, err := RS.sserial.Read(buf)
	if err != nil {
		panic(err)
	}
	if len == 0 {
		return nil, errors.New("READ empty byte")
	}
	return buf[:len], nil

}
func (RS *cserial) Flush() error {
	err := RS.sserial.Flush()
	return err
}
func (RS *cserial) Disconn() {
	RS.sserial.Close()
}

type netWork struct {
	name     string
	ip       string
	port     int
	connnect *net.TCPConn
}

func (network *netWork) Conn() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", network.ip)
	if err != nil {
		panic(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		panic(err)
	}
	network.connnect = conn
	defer network.Disconn()
}
func (network *netWork) Write(b []byte) {
	network.connnect.Write(b)
}
func (network *netWork) Read() ([]byte, error) {
	return nil, nil
}
func (network *netWork) Flush() error {
	return errors.New("flush")
}
func (network *netWork) Disconn() {
	network.connnect.Close()
}
func NewNetWork(name string, ip string, port int, protocol string) *netWork {
	return &netWork{
		name: name,
		ip:   ip,
		port: port,
	}
}
