package main

import "fmt"

func main() {

	number := 5
	fmt.Println("1. Alamat Memory awal", &number)
	fmt.Println("2. Nilai Awal", number)

	change(&number, 100)

	fmt.Println("5. Nilai Akhir", number)
	fmt.Println("6. Alamat Memory akhir ", &number)

}

func change(old *int, new int) {
	*old = new
	fmt.Println("3. Nilai Akhir", &old)
	fmt.Println("4. in function =", *old)

}
