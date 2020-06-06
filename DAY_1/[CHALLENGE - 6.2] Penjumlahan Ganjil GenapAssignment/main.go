// Penjumlahan Ganjil Genap
package main

import (
	"fmt"
)

func main() {
	var angka int
	var jumlah int
	jumlah = 0
	fmt.Printf("Masukan Angka : ")
	fmt.Scan(&angka)
	for i := 1; i < angka; i++ {
		if i%2 != 0 {
			fmt.Print(i, " + ")
			jumlah = jumlah + i
		}

	}
	fmt.Print("= ", jumlah)
}
