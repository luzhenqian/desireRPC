package client

import (
	"desire/serializer"
)

// Client rpc client 接口
type Client interface {
	Call(name string, params ...interface{}) (interface{}, error)
	RegisterType(types ...serializer.Type)
}
