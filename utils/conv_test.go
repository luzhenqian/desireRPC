package utils

import "testing"

func TestUint16ToBytes(t *testing.T) {
	var version uint16 = 12
	bytes := Uint16ToBytes(version)
	t.Log(bytes)
}

func TestStringToBytes(t *testing.T) {
	messageType := "a"
	t.Log([]byte(messageType))
}
