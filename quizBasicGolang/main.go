package main

import "fmt"

func main() {
	title := "Golang The best Lenguage"

	for index, letter := range title {
		if index%2 == 0 {
			fmt.Println("cetak index genap :", index, "letter: ", string(letter))
			// fmt.Println("index :", index, " Letter :", string(letter))
		}
	}

	fmt.Println("----------------------------------------------------------------")

	for index, letter := range title {
		kata := string(letter)
		// if string(letter) == "a" || string(letter) == "i" || string(letter) == "u" || string(letter) == "e" || string(letter) == "o" {
		// 	fmt.Println("cetak index vocal :", index, "letter: ", string(letter))
		// 	// fmt.Println("index :", index, " Letter :", string(letter))
		// }
		// fmt.Println("----------------------------------------------------------------")

		switch kata {
		case "a", "i", "u", "e", "o":
			fmt.Println("cetak index vocal :", index, "letter: ", string(letter))
		}
	}

}
