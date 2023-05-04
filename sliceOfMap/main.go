package main

import "fmt"

func main() {
	students := []map[string]string{
		{"name": "alwani", "score": "A"},
		{"name": "love", "score": "C"},
		{"name": "nurul", "score": "A"},
	}

	// fmt.Println(students)

	for _, key := range students {
		fmt.Println("Name:", key["name"], " || ", "Score:", key["score"])
	}
}
