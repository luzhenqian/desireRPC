package main

import (
	"desire/client"
	"fmt"
)

type User struct {
	ID   uint8
	Name string
	Age  uint8
}

type UserQueryParam struct {
	ID uint8
}

func main() {
	services := client.Services{
		"UserServices": "127.0.0.1:8888",
	}
	clt := client.New(services)
	clt.RegisterType(UserQueryParam{})

	user := &User{}
	err := clt.Call("UserServices", client.CallOption{
		FunctionName: "getUser",
		RequestData:  UserQueryParam{ID: 1},
		ResponseData: user,
	})
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("user:", user)
}
