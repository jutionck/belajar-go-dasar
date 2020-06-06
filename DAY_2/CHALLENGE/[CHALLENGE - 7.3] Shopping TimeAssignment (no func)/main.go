package main

import "fmt"

func main() {
	var barang = []int{15000, 12000, 5000, 3000, 10000}
	var saldo = 30000
	HowManyGifts(saldo, barang)
}

func HowManyGifts(saldo int, barang []int) {
	for i := 0; i < len(barang)-1; i++ {
		for j := i + 1; j < len(barang); j++ {
			// urut manual
			if barang[i] > barang[j] {
				sementara := barang[i]
				barang[i] = barang[j]
				barang[j] = sementara
				// fmt.Println(barang[i], barang[j])
			}
		}
	}
	jumlahbeli := 0 //isi nilai awal 0
	for k := 0; k <= len(barang)-1; k++ {
		if barang[k] <= saldo {
			saldo -= barang[k] //ex : 10000 = 10000 - gift yang index of nya dari loop k
			jumlahbeli++       //akan terus bertambah 1 sampai budgetnya abis
		}
	}
	fmt.Println(jumlahbeli, "Barang")

}
