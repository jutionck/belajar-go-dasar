// Penjumlahan Pasangan
package main

import "fmt"

func main() {
	var deret1 = []int{1, 2, 3, 4, 5}
	var sum int = 8
	fmt.Println(PenjumalahPasangan(deret1, sum))
}

func PenjumalahPasangan(angka []int, total int) (hitung int) {
	// var temp string
	for i := 0; i < len(angka); i++ {
		for j := i; j < len(angka); j++ {
			if angka[i]+angka[j] == total {
				temp := fmt.Sprintf("[[%d %d]]", angka[i], angka[j])
				hitung = len(temp)
			}
		}
	}
	return
}
