// Deret kelipatan konsisten
// Jution Candra Kirana

package main

import "fmt"

// func is chek number of deret
func deretCheck(number []int) bool {
	difference := findDifference(number)
	var isCheck bool = false //return
	for i := 1; i <= len(number)-1; i++ {
		for j := 0; j < len(number); j++ {
			if i == j+1 {
				if number[i]-number[j] == difference {
					isCheck = true
				} else {
					isCheck = false
				}
			}
		}
	}
	return isCheck
}

// find a difference
func findDifference(number []int) int {
	var difference int //return
	if number[1]-number[0] == number[2]-number[1] {
		difference = number[1] - number[0]
	}
	return difference
}

func main() {
	var deretOfNumber = []int{2, 4, 6, 8}
	idCheck := deretCheck(deretOfNumber)
	if idCheck {
		fmt.Println(idCheck)
	} else {
		fmt.Println(idCheck)
	}
}
