package main

import (
	"fmt"
)

func main() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("elemen %d : %s\n", i, fruits[i])
	// }

	// for j := len(fruits) - 1; j > 0; j-- {
	// 	fmt.Printf("elemen %d : %s\n", j, fruits[j])
	// }

	var i = len(fruits) - 1
	for i >= 0 {
		fmt.Println(i)
		i--
	}
}
