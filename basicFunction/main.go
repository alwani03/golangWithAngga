package main

func main() {
	retun1 := printString("Golang")
	retun2 := printString("Php")
	retun3 := printString("Javascript")

	println(retun1)
	println(retun2)
	println(retun3)

}

func printString(sentence string) string {
	sentence = "I Learn " + sentence
	return sentence
}
