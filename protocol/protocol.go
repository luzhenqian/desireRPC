package protocol

// Protocol 协议
type Protocol interface {
	Pack() ([]byte, error)
	Unpack(b []byte, payload interface{}) error
}
