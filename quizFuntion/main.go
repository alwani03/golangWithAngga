package main

import (
	"errors"
	"fmt"
)

func main() {
	scores := []int{10, 5, 8, 9, 7}
	totalScore := sum(scores)
	fmt.Println(totalScore)

	result, err := calculate(10, 2, "m")

	// fmt.Println(result)
	// fmt.Println(err)

	if err != nil {
		fmt.Println("Terjadi Kesalahan")
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}

}

func sum(array []int) int {
	var total int
	for _, v := range array {
		total += v
	}
	return total
}

func calculate(number1, number2 int, operation string) (int, error) {

	var result int
	var errorResult error

	if operation == "+" {
		result = number1 + number2
	} else if operation == "-" {
		result = number1 - number2
	} else if operation == "*" {
		result = number1 - number2
	} else if operation == "/" {
		result = number1 - number2
	} else {
		errorResult = errors.New("Operation Not Found")
	}
	return result, errorResult
}
