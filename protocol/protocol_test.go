package protocol

import (
	"testing"
)

type User struct {
	Name string
	Age  int
}

func TestPack(t *testing.T) {
	d := Desire{
		Payload: User{
			"洋子", 25,
		},
	}
	b, err := d.Pack()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(b)

	u := &User{}
	dp := Desire{}
	err = dp.Unpack(b, u)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dp, dp.Payload)
}
