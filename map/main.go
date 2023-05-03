package main

import "fmt"

func main() {

	var myMap map[string]int
	myMap = map[string]int{}

	myMap["go"] = 12
	myMap["js"] = 12
	myMap["c"] = 12

	fmt.Println(myMap)
	fmt.Println(myMap["go"])

	fmt.Println("----------")

	mymApSimple := map[string]string{
		"food":    "rice",
		"drink":   "tea",
		"dissert": "cake",
	}

	fmt.Println(mymApSimple)
	fmt.Println(mymApSimple["food"])

}
