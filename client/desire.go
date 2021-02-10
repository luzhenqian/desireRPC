package client

import (
	"desire/protocol"
	"desire/serializer"
	"net"
	"time"
)

// Desire client 实现
type Desire struct {
	Services Services
	pool     ConnectionPool
	p        protocol.Protocol
	s        serializer.Serializer
}

type Options func()

func New(services Services, options ...Options) *Desire {
	// TODO
	return &Desire{Services: services, pool: make(ConnectionPool)}
}

// func New(p protocol.Protocol, s serializer.Serializer) *Desire {
// 	return &Desire{}
// }

// Address rpc 服务端地址
type Address string

// ConnectionPool tcp 连接池
type ConnectionPool map[Address]net.Conn

// Option 配置项
type Option struct {
	KeepAlive    time.Duration
	Header       protocol.VariableHeader
	FunctionName FunctionName
	RequestData  RequestData
	ResponseData ResponseData
}

type FunctionName string
type RequestData interface{}
type ResponseData interface{}

type Services map[ServiceName]Address
type ServiceName string

func (d *Desire) Call(serviceName ServiceName, option Option) error {
	if addr, ok := d.Services[serviceName]; ok {
		return d.call(addr, option)
	}
	// TODO error
	return nil
}

func (d *Desire) RegisterType(types ...serializer.Type) {
	d.s.RegisterType(types)
}

// Call 调用 RPC 服务
func (d *Desire) call(addr Address, option Option) error {
	conn, err := d.getConn(addr)
	if err != nil {
		return err
	}
	err = d.write(conn, option)
	if err != nil {
		return err
	}
	err = d.read(conn, option.ResponseData)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

func (d *Desire) getConn(addr Address) (net.Conn, error) {
	if conn, ok := d.pool[addr]; ok {
		if ok {
			return conn, nil
		}
	}
	conn, err := net.Dial("tcp", string(addr)) // 发起 RPC
	d.pool[addr] = conn
	// TODO keepalive
	return conn, err
}
