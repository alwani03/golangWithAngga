package main

import "fmt"

func main() {
	broad, arround := calculate(12, 12)

	fmt.Println(broad, "||", arround)
}

func calculate(length, width int) (int, int) {
	broad := length * width
	arround := 2 * (length * width)

	return broad, arround
}
