package main

import (
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func main() {

	user := User{1,
		"I",
		"Love",
		"You",
		true,
	}

	user2 := User{2,
		"Nurul",
		"Karomah",
		"nurul@gmail.com",
		true,
	}

	fmt.Println("user -->", user)
	fmt.Println("user2 -->", user2)

	displayUser1 := displayUser(user)
	displayUser2 := displayUser(user2)

	fmt.Println("displayUser -->", displayUser1)
	fmt.Println("displayUser2 -->", displayUser2)

}

func displayUser(user User) string {
	return fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
}
