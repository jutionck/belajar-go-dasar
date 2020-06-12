// Shooping Time
package main

import (
	"fmt"
	"sort"
)

func main() {
	var barang = []int{1500000, 500000, 250000, 175000, 50000}
	var saldo = 700000
	HowManyGifts(saldo, barang)
}

func HowManyGifts(saldo int, barang []int) {
	// pengurutan barang dari mahal
	sort.Slice(barang, func(i, j int) bool {
		return barang[i] > barang[j]
	})

	jumlahbeli := 0 //isi nilai awal 0
	for k := 0; k <= len(barang)-1; k++ {
		if barang[k] <= saldo {
			saldo -= barang[k] //ex : 10000 = 10000 - gift yang index of nya dari loop k
			jumlahbeli++       //akan terus bertambah 1 sampai budgetnya abis
		}
	}
	// Sisa saldo
	fmt.Println(saldo)
	fmt.Println(jumlahbeli, "Barang")

}
