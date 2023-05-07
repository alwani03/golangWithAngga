package main

import (
	"fmt"
	"strconv"
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

func (group Group) Groups() (name string,
	totalAnggota string,
	nameAnggota []string,
) {
	name = "Name Group: " + group.Name
	totalAnggota = "Jumlah Anggota: " + strconv.Itoa(len(group.Users))
	for _, value := range group.Users {
		nameAnggota = append(nameAnggota, value.FirstName)

	}

	return
}

func main() {

	user := User{1,
		"I",
		"Love",
		"You",
		true,
	}

	user2 := User{1,
		"Me",
		"Love",
		"You to",
		true,
	}

	whoUser := []User{user, user2}
	data := Group{"Gamer", user, whoUser, true}
	nameGroup, totalAnggota, nameAnggota := data.Groups()

	fmt.Println(nameGroup)
	fmt.Println(totalAnggota)

	for _, v := range nameAnggota {
		fmt.Println(v)
	}

}
