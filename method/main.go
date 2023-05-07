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

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvalaible bool
}

func (user User) Display() string {

	return fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)

}

func main() {

	user := User{1,
		"I",
		"Love",
		"You",
		true,
	}

	user2 := User{1,
		"I",
		"Love",
		"You to",
		true,
	}

	result := user.Display()
	result2 := user2.Display()

	fmt.Println(result)
	fmt.Println(result2)

}

func displayGroup(grup Group) {
	fmt.Printf("Name Group: %s", grup.Name)
	fmt.Println("")
	fmt.Printf("Jumlah Anggota: %d", len(grup.Users))
	fmt.Println("")

	for _, value := range grup.Users {
		fmt.Println(value.FirstName)
	}
}
