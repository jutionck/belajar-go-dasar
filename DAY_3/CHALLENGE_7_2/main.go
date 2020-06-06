// Nilai Terbesar dan Terkecil
package main

import "fmt"

func main() {
	var isideret = []int{1, 2, 5}
	max, min := CalculatorStatic(isideret)
	fmt.Printf("Nilai Terbesar = %d\n", max)
	fmt.Printf("Nilai Terkecil = %d\n", min)
}

func CalculatorStatic(deret []int) (int, int) {
	var max, min int
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
	return max, min
}
