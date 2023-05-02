package main

import (
	"fmt"
	"pertama/calculation"
	"pertama/perkalian"
)

func main() {
	fmt.Println("hallo, Belajar Golang")

	fmt.Println("--------")
	result := calculation.Add(5, 5)
	fmt.Println(result)
	fmt.Println("--------")

	fmt.Println("hallo, Ini Kuis")
	hasil := perkalian.Kali(5, 10)
	fmt.Println(hasil)

}
