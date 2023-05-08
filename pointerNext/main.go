package main

import "fmt"

func main() {

	var numberA int = 5
	var numberB *int = &numberA

	fmt.Println(numberA)
	fmt.Println(numberB)
	fmt.Println(*numberB)

	numberA = 20

	fmt.Println(numberA)
	fmt.Println(numberB)
	fmt.Println(*numberB)

}
