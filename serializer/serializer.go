package serializer

import "desire/types"

type Type interface{}

type Serializer interface {
	RegisterType(types ...Type)
	Marshal(data Type) (raw []byte, err error)
	UnMarshal(raw []byte, data Type) error

	MarshalLength(len types.Length) []byte
	MarshalHeaderLength(headerLen types.HeaderLength) []byte
	MarshalVersion(version types.Version) []byte
	MarshalMessageID(msgID types.MessageID) []byte

	UnmarshalHeaderLength([]byte) types.HeaderLength
}
