package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func main() {

	user := User{}

	user.ID = 3
	user.FirstName = "Nurul"
	user.LastName = "Karomah"
	user.Email = "Nurul@gmail.com"
	user.IsActive = true

	user2 := User{ID: 2,
		FirstName: "Achmad",
		LastName:  "Alwani",
		Email:     "Alwani@gmail.com",
		IsActive:  true,
	}

	//or
	user3 := User{1,
		"I",
		"Love",
		"You",
		true,
	}

	fmt.Println(user)
	fmt.Println(user2)
	fmt.Println(user3)

}
