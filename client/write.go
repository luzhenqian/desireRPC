package client

import (
	"desire/protocol"
	"net"
)

func (d *Desire) write(conn net.Conn, option Option) error {
	ptcBytes, err := makeProtocolBytes(option.FunctionName, option.RequestData, option.Header)
	if err != nil {
		return err
	}
	_, err = conn.Write(ptcBytes)
	return err
}

type Request struct {
	FunctionName FunctionName
	RequestData  RequestData
}

func makeProtocolBytes(fnName FunctionName, data RequestData, header protocol.VariableHeader) ([]byte, error) {
	payload := Request{
		FunctionName: fnName,
		RequestData:  data,
	}
	ptc := protocol.Desire{
		Payload: payload,
	}
	ptc.Header.Variable = header
	return ptc.Pack()
}
