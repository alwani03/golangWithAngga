package main

import "fmt"

func main() {
	var totalValue float32
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	for _, v := range scores {
		totalValue = totalValue + float32(v)
	}

	total := len(scores)
	rataRata := totalValue / float32(total)
	fmt.Println("Rata-rata dari Score tersebut adalah", rataRata)
	fmt.Println("===============================================")

	var goodScore []int

	for _, c := range scores {
		if c >= 90 {
			goodScore = append(goodScore, c)
		}
	}

	fmt.Println("Slice :", goodScore)

}
