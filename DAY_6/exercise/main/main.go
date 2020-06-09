// import "exercise/formula"
package main

import (
	"fmt"
)

var penjumlahan = func(a, b int) int {
	return a + b
}

func main() {
	angka1 := 10
	angka2 := 5
	func() { //IIFE function
		var jumlah int
		jumlah = angka1 + angka2
		fmt.Println(jumlah)
	}() //parameter
}

func cetak(nilai func(int, int) int, a, b int) {
	fmt.Println(nilai(a, b))
}
