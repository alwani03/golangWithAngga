package main

import "fmt"

func main() {
	broad, arrounds := calculate(12, 12)

	fmt.Println(broad, "||", arrounds)
}

func calculate(length, width int) (broad int, arround int) {
	broad = length * width
	arround = 2 * (length * width)

	return
}
