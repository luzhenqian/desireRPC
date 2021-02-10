package utils

import (
	"bytes"
	"encoding/binary"
	"reflect"
)

func IsZeroOfUnderlyingType(data interface{}) bool {
	return reflect.DeepEqual(data, reflect.Zero(reflect.TypeOf(data)).Interface())
}

func StringToBytes(s string) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, s)
	return bytesBuffer.Bytes()
}

func Uint8ToBytes(i uint8) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, i)
	return bytesBuffer.Bytes()
}

func Uint16ToBytes(i uint16) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, i)
	return bytesBuffer.Bytes()
}

func Uint32ToBytes(i uint32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, i)
	return bytesBuffer.Bytes()
}

func Float32ToBytes(i float32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, i)
	return bytesBuffer.Bytes()
}

func BytesToUint8(b []byte) uint8 {
	bytesBuffer := bytes.NewBuffer(b)
	var ret uint8
	binary.Read(bytesBuffer, binary.BigEndian, &ret)
	return ret
}

func BytesToUint16(b []byte) uint16 {
	bytesBuffer := bytes.NewBuffer(b)
	var ret uint16
	binary.Read(bytesBuffer, binary.BigEndian, &ret)
	return ret
}

func BytesToUint32(b []byte) uint32 {
	bytesBuffer := bytes.NewBuffer(b)
	var ret uint32
	binary.Read(bytesBuffer, binary.BigEndian, &ret)
	return ret
}
