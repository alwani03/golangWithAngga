package main

import (
	"defExpoAndUnexpo/management"
	"fmt"
)

/*Keterangan kalo huruf kecil tidak bisa dipanggil dibeda package*/

func main() {

	user := management.User{1,
		"I",
		"Love",
		"You",
		true,
	}

	user2 := management.User{1,
		"Me",
		"Love",
		"You to",
		true,
	}

	result := user.Display()
	result2 := user2.Display()

	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println("================")

	whoUser := []management.User{user, user2}
	groups := management.Group{"Gamer", user, whoUser, true}

	groups.DisplayGroup()

}
