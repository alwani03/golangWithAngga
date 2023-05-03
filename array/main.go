package main

import "fmt"

func main() {
	var leuanges [5]string
	leuanges[0] = "Go"
	leuanges[1] = "Ruby"
	leuanges[2] = "php"
	leuanges[3] = "javascript"
	leuanges[4] = "c++"

	fmt.Println(leuanges)
	length := len(leuanges)
	fmt.Println(length)

	leuangesGolang := [...]string{
		"Go",
		"Ruby",
		"php",
		"javascript",
		"c++",
	}

	fmt.Println(leuangesGolang)

	for index, lang := range leuangesGolang {
		fmt.Println("---> Index :", index, "lang :", lang)
	}

}
