package main

import "fmt"

func main() {
	age := 10

	if age > 10 {
		fmt.Println("kena")
	} else {
		fmt.Println("ga kena")
	}

	score := 80
	var grade string

	if score <= 10 {
		grade = "E"
	} else if score <= 60 {
		grade = "C"
	} else {
		grade = "Bagus"
	}

	println(grade)
}
