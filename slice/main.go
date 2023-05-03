package main

import "fmt"

func main() {
	var gamingConsole []string

	gamingConsole = append(gamingConsole, "Ps8")
	gamingConsole = append(gamingConsole, "xbox")
	gamingConsole = append(gamingConsole, "gambot")
	fmt.Println(gamingConsole)

	games := []string{"marioBros", "GTA5", "bola"}
	fmt.Println(games)

	fmt.Println("----------")

	for _, v := range games {
		fmt.Println(v)
	}

}
