/*!
 * <------------------------ MODBUS SERIAL LINE ADU (1) ------------------->
 *              <----------- MODBUS PDU (1') ---------------->
 *  +-----------+---------------+----------------------------+-------------+
 *  | Address   | Function Code | Data                       | CRC/LRC     |
 *  +-----------+---------------+----------------------------+-------------+
 *  |           |               |                                   |
 * (2)        (3/2')           (3')                                (4)
 */
package modbus

import (
	"bytes"
	"encoding/binary"
	"errors"
	"reflect"
)

// Modbus protocol behavior
type Modbus interface {
	SetCmd(raw []byte)
	GetCmd() []byte
	Unpack(data []byte, kind reflect.Kind) (interface{}, error)
}

// RtuCmd struct is modbusRTU
type rtuCmd struct {
	Address  byte
	FuncCode byte
	Data     []byte
	Crc      [2]byte
}

// SetCmd for add crc16
func (cmd *rtuCmd) SetCmd(raw []byte) {
	addr := raw[0]
	fn := raw[1]
	if len(raw[2:]) < 4 {
		panic("modbus code needs more bytes")
	}
	data := raw[2:]
	cmd.Address = addr
	cmd.FuncCode = fn
	cmd.Data = data
	crc := makeCrc(raw)
	cmd.Crc = crc
}

// GetCmd return full modbus code within crc16
func (cmd *rtuCmd) GetCmd() []byte {

	conbineCmd := []byte{cmd.Address, cmd.FuncCode}
	conbineCmd = append(conbineCmd, cmd.Data...)
	for _, v := range cmd.Crc {
		conbineCmd = append(conbineCmd, v)
	}
	return conbineCmd

}

// Unpack resolve byte array to int or float warpper
func (cmd *rtuCmd) Unpack(data []byte, kind reflect.Kind) (interface{}, error) {
	if len(data) < 1 {
		return nil, errors.New("can't Unpacked")
	}
	decodeNumber := data[2]
	decodeByte := data[3:(3 + decodeNumber)]
	switch kind {
	case reflect.Int8:
		return int8(decodeByte[0]), nil
	case reflect.Int16:
		return int16(binary.BigEndian.Uint16(decodeByte)), nil
	case reflect.Int32:
		return int32(binary.BigEndian.Uint32(decodeByte)), nil
	case reflect.Int64:
		return int64(binary.BigEndian.Uint64(decodeByte)), nil
	case reflect.Uint8:
		return decodeByte[0], nil
	case reflect.Uint16:
		return binary.BigEndian.Uint16(decodeByte), nil
	case reflect.Uint32:
		return binary.BigEndian.Uint32(decodeByte), nil
	case reflect.Uint64:
		return binary.BigEndian.Uint64(decodeByte), nil
	case reflect.Float32:
		var v float32
		buf := bytes.NewBuffer(decodeByte)
		binary.Read(buf, binary.BigEndian, &v)
		return v, nil
	case reflect.Float64:
		var v float64
		buf := bytes.NewBuffer(decodeByte)
		binary.Read(buf, binary.BigEndian, &v)
		return v, nil
	default:
		return nil, errors.New("unknow type")
	}
}
func makeCrc(data []byte) [2]byte {
	crc := 0xFFFF
	for _, v := range data {
		crc ^= int(v)
		for i := 0; i < 8; i++ {
			if (crc & 1) != 0 {
				crc >>= 1
				crc ^= 0xA001
			} else {
				crc >>= 1
			}
		}
	}
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, uint16(crc))
	result := [2]byte{}
	copy(result[:], buf)
	return result
}

// NewRtuCmd return *RtuCmd
func NewRtuCmd() *rtuCmd {
	return &rtuCmd{}
}

type tcpCmd struct {
	*tcpHeader
	Address  byte
	FuncCode byte
	Data     []byte
}
type tcpHeader struct {
	transactionID [2]byte
	protocolID    [2]byte
	length        [2]byte
}

func (cmd *tcpCmd) SetCmd(addr byte, fn byte, data []byte) {
	cmd.tcpHeader = &tcpHeader{
		transactionID: [2]byte{0x00, 0x00},
		protocolID:    [2]byte{0x00, 0x00},
		length:        [2]byte{0x00, 0x00},
	}
	cmd.Address = addr
	cmd.FuncCode = fn
	cmd.Data = data
}
func (cmd *tcpCmd) GetCmd() []byte {

	combineCmd := []byte{}
	for _, v := range cmd.tcpHeader.transactionID {
		combineCmd = append(combineCmd, v)
	}
	for _, v := range cmd.tcpHeader.protocolID {
		combineCmd = append(combineCmd, v)
	}
	for _, v := range cmd.tcpHeader.length {
		combineCmd = append(combineCmd, v)
	}
	combineCmd = append(combineCmd, cmd.Address)
	combineCmd = append(combineCmd, cmd.FuncCode)
	combineCmd = append(combineCmd, cmd.Data...)

	return combineCmd

}
func (cmd *tcpCmd) Unpack(data []byte, kind reflect.Kind) (interface{}, error) {
	return nil, nil
}
func NewTcpCmd() *tcpCmd {
	return &tcpCmd{}
}
