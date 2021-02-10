package server

import (
	"fmt"
	"testing"
)

func TestRegister(t *testing.T) {
	// 服务端代码
	desireServer := New()
	err := desireServer.Register(add)
	if err != nil {
		fmt.Println("Register err:", err)
	}

	// 客户端代码
	// res, err := desireServer.Call("add", 1, 2)
	// if err != nil {
	// 	fmt.Println("err:", err)
	// }
	// fmt.Println("res:", res)
}

func add(a int, b int) int {
	return a + b
}
