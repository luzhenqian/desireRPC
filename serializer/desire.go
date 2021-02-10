package serializer

import (
	"bytes"
	"desire/utils"
	"encoding/gob"
	"io"
)

type Desire struct{}

func (d *Desire) RegisterType(types ...Type) {
	for _, t := range types {
		gob.Register(t)
	}
}

func (d *Desire) Serializer(data Type) ([]byte, error) {
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

func (d *Desire) Unserializer(raw []byte, data Type) error {
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
