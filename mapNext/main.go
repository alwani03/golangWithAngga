package main

import "fmt"

func main() {

	mymApSimple := map[string]string{
		"food":    "rice",
		"drink":   "tea",
		"dissert": "cake",
	}

	for key, v := range mymApSimple {
		fmt.Println("this key:", key, "||", "this value:", v)
	}

	fmt.Println("========================")

	delete(mymApSimple, "dissert")

	for key, v := range mymApSimple {
		fmt.Println("this key:", key, "||", "this value:", v)
	}

	fmt.Println("========================")

	value, check := mymApSimple["fruit"]
	fmt.Println(value)
	fmt.Println(check)

}
