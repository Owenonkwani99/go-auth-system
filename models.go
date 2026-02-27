package main

import "fmt"

type User struct {
	Username string
	Password string
}

var demoUser = User{
	Username: "Owen",
	Password: "Double007",
} 

func (u User) String() string {
	return fmt.Sprintf("User{Username: %s}", u.Username)
}
