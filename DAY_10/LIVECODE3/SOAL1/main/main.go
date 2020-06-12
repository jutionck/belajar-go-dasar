// Program Penjumlahan Pasangan
// Output - Hitung Jumlah nya
// By Jution Candra Kirana

package main

import (
	"LIVECODE3/SOAL1/service"
	"fmt"
)

func main() {
	var listOfNumber = []int{1, 2, 3, 4, 5}
	var sumNumber int = 8
	pairsNumber := service.SliceNumber{listOfNumber, sumNumber}
	printResult(pairsNumber)
}

func printResult(p service.NumberResult) {
	result := p.PairsNumber()
	fmt.Println("Jumlah Penjumlahan Pasangan adalah :", result)
}
