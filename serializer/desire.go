package serializer

import (
	"bytes"
	"desire/types"
	"desire/utils"
	"encoding/binary"
	"encoding/gob"
	"io"
	"reflect"
)

type Desire struct{}

func (d *Desire) RegisterType(types ...Type) {
	for _, t := range types {
		gob.Register(t)
	}
}

func (d *Desire) Marshal(data Type) ([]byte, error) {
	if data == nil || utils.IsZeroOfUnderlyingType(data) {
		return nil, nil
	}
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (d *Desire) UnMarshal(raw []byte, data Type) error {
	if len(raw) == 0 {
		return nil
	}
	buf := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(data)
	if err != nil {
		if err != io.EOF {
			return err
		}
	}
	return nil
}

func (d *Desire) MarshalLength(len types.Length) []byte {
	return uint32ToBytes(uint32(len))
}

func (d *Desire) MarshalHeaderLength(headerLen types.HeaderLength) []byte {
	return uint16ToBytes(uint16(headerLen))
}

func (d *Desire) MarshalVersion(version types.Version) []byte {
	return uint16ToBytes(uint16(version))
}

func (d *Desire) MarshalMessageID(msgID types.MessageID) []byte {
	return uint16ToBytes(uint16(msgID))
}

func (d *Desire) UnmarshalHeaderLength(b []byte) types.HeaderLength {
	return types.HeaderLength(bytesToUint16(b))
}

func IsZeroOfUnderlyingType(data interface{}) bool {
	return reflect.DeepEqual(data, reflect.Zero(reflect.TypeOf(data)).Interface())
}

func stringToBytes(s string) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, s)
	return bytesBuffer.Bytes()
}

func uint8ToBytes(i uint8) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, i)
	return bytesBuffer.Bytes()
}

func uint16ToBytes(i uint16) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, i)
	return bytesBuffer.Bytes()
}

func uint32ToBytes(i uint32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, i)
	return bytesBuffer.Bytes()
}

func float32ToBytes(i float32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, i)
	return bytesBuffer.Bytes()
}

func bytesToUint8(b []byte) uint8 {
	bytesBuffer := bytes.NewBuffer(b)
	var ret uint8
	binary.Read(bytesBuffer, binary.BigEndian, &ret)
	return ret
}

func bytesToUint16(b []byte) uint16 {
	bytesBuffer := bytes.NewBuffer(b)
	var ret uint16
	binary.Read(bytesBuffer, binary.BigEndian, &ret)
	return ret
}

func bytesToUint32(b []byte) uint32 {
	bytesBuffer := bytes.NewBuffer(b)
	var ret uint32
	binary.Read(bytesBuffer, binary.BigEndian, &ret)
	return ret
}
