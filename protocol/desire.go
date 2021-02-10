package protocol

import (
	"desire/serializer"
	"desire/utils"
)

// Desire 协议
type Desire struct {
	Header  Header
	Payload interface{}
	s       serializer.Serializer
}

// Header 协议头
type Header struct {
	FixedHeader                // 10 byte
	Variable    VariableHeader // n byte
}

// FixedHeaderLen 固定头长度
const FixedHeaderLen = 10

// LengthLen 长度字段长度
const LengthLen = 4

// FixedHeader 固定协议头
type FixedHeader struct {
	Length       Length // 0 1 2 3 byte
	HeaderLength uint16 // 4 5 byte
	Version      uint16 // 6 7 byte
	MessageID    uint16 // 8 9 byte
}

// Length 数据包整体长度
type Length uint32

// VariableHeader 可变协议头
type VariableHeader map[string]string

// Pack 打包
func (d *Desire) Pack() ([]byte, error) {
	var ret []byte

	// 协议体
	payload, err := d.s.Serializer(d.Payload)

	if err != nil {
		return nil, err
	}
	payloadLength := len(payload)

	// 扩展协议头
	variable, err := d.s.Serializer(d.Header.Variable)
	if err != nil {
		return nil, err
	}
	variableLength := len(variable)
	headerLength := 10 + variableLength

	length := headerLength + payloadLength

	// TODO 消除 utils 转到 serializer
	lengthBytes := utils.Uint32ToBytes(uint32(length))
	ret = append(ret, lengthBytes...) // 设置完整报文长度

	headerLengthBytes := utils.Uint16ToBytes(uint16(headerLength))
	ret = append(ret, headerLengthBytes...) // 设置协议头长度

	versionBytes := utils.Uint16ToBytes(d.Header.Version) // 版本
	ret = append(ret, versionBytes...)                    // 设置协议版本

	messageIDBytes := utils.Uint16ToBytes(d.Header.MessageID)
	ret = append(ret, messageIDBytes...) // 设置消息 ID

	ret = append(ret, payload...) // 设置协议体

	return ret, nil
}

// Unpack 拆包
// 第二个参数 payload 需要是一个指针类型，拆包过程会给 payload 赋值
func (d *Desire) Unpack(b []byte, payload interface{}) error {
	d.Header.Length = Length(utils.BytesToUint32(b[:LengthLen]))
	d.Header.HeaderLength = utils.BytesToUint16(b[LengthLen : FixedHeaderLen-LengthLen])
	payloadBytes := b[d.Header.HeaderLength:]
	err := d.s.Unserializer(payloadBytes, payload)
	if err != nil {
		return err
	}
	d.Payload = payload
	return nil
}
