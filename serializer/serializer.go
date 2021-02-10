package serializer

type Type interface{}

type Serializer interface {
	RegisterType(types ...Type)
	Serializer(data Type) (raw []byte, err error)
	Unserializer(raw []byte, data Type) error
}
