package main

import "fmt"

func main() {
	var isideret = []int{1, 2, 5}
	fmt.Println(CalculatorStatic(isideret))
}

func CalculatorStatic(deret []int) (int, int, int) {
	var max, min, total, ratarata int
	max = deret[0]
	min = deret[0]
	// max min
	for i := 1; i < len(deret); i++ {
		if deret[i] > max {
			max = deret[i]
		}
		if deret[i] < min {
			min = deret[i]
		}
	}
	// rata rata
	total = 0
	for i := 0; i < len(deret); i++ {
		total += deret[i]
	}
	ratarata = total / len(deret)
	return max, min, ratarata
}
