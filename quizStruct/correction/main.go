package main

import "fmt"

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

func (groups Group) DisplayGroup() {
	fmt.Printf("Name Group: %s", groups.Name)
	fmt.Println("")
	fmt.Printf("Jumlah Anggota: %d", len(groups.Users))
	fmt.Println("")

	for _, value := range groups.Users {
		fmt.Println(value.FirstName)
	}
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

	whoUser := []User{user, user2}

	groups2 := Group{"Gamer", user, whoUser, true}

	groups2.DisplayGroup()

}
