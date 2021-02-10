package protocol

import (
	"bytes"
	"encoding/binary"
	"net"
)

// Response 返回数据
type Response interface{}

func (d *Desire) Read(conn net.Conn, response Response) error {
	length, err := readLength(conn)
	if err != nil {
		return err
	}
	bs, err := readAll(conn, length)
	if err != nil {
		return err
	}
	dp := Desire{}
	return dp.Unpack(bs, response)
}

func readLength(conn net.Conn) (Length, error) {
	b := make([]byte, LengthLen)
	if _, err := conn.Read(b); err != nil {
		return 0, err
	}
	bytesBuffer := bytes.NewBuffer(b)
	var length Length
	binary.Read(bytesBuffer, binary.BigEndian, &length)
	return length, nil
}

func readAll(conn net.Conn, length Length) ([]byte, error) {
	b := make([]byte, int(length-LengthLen))
	if _, err := conn.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}
