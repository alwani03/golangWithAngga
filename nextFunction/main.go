package main

import "fmt"

func main() {
	summation := add(10, 10)

	fmt.Println(summation)
}

func add(number, number2 int) int {

	return number + number2

}
