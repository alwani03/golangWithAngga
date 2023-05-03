package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("golang ke : ", i)
	}

	fmt.Println("-----------")

	j := 1
	for j <= 10 {
		fmt.Println("golang while ke : ", j)
		j++
	}

	title := "Golang The best Lenguage"

	for index, letter := range title {
		fmt.Println("index :", index, " Letter :", string(letter))
	}

}
