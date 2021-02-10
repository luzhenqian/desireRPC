package main

import (
	"desire/server"
)

func main() {
	sev := server.New()
	sev.Register(getUser)
	sev.Run(":8888")
}

func getUser(p UserQueryParam) *User {
	for _, user := range users {
		if user.ID == p.ID {
			return &user
		}
	}
	return nil
}

var users []User = []User{
	{
		ID:   1,
		Name: "q",
		Age:  35,
	},
	{
		ID:   2,
		Name: "h",
		Age:  36,
	},
}

type User struct {
	ID   uint8
	Name string
	Age  uint8
}

type UserQueryParam struct {
	ID uint8
}
