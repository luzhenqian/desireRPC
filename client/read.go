package client

import (
	"bytes"
	"desire/protocol"
	"encoding/binary"
	"net"
)

func (d *Desire) read(conn net.Conn, response interface{}) error {
	length, err := readLength(conn)
	if err != nil {
		return err
	}
	bs, err := readAll(conn, length)
	if err != nil {
		return err
	}
	dp := protocol.Desire{}
	return dp.Unpack(bs, response)
}

func readLength(conn net.Conn) (protocol.Length, error) {
	b := make([]byte, protocol.LengthLen)
	if _, err := conn.Read(b); err != nil {
		return 0, err
	}
	bytesBuffer := bytes.NewBuffer(b)
	var length protocol.Length
	binary.Read(bytesBuffer, binary.BigEndian, &length)
	return length, nil
}

func readAll(conn net.Conn, length protocol.Length) ([]byte, error) {
	b := make([]byte, int(length-protocol.LengthLen))
	if _, err := conn.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}
