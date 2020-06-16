/*
Enigma Sale
Perhatikan uang yang kamu punya, kira kira dengan uang yang kamu punya dapat membeli barang apa saja. Berikut kriteria yang harus di perhatikan :
1. Jika id kosong maka tampilkan "Mohon maaf, toko Enigma Jaya hanya berlaku untuk member saja"
2. Jika uang yang dimiliki kurang dari 50000 maka tampilkan "Mohon maaf, uang tidak cukup"
3. Member yang berbelanja di toko Enigma Jaya akan membeli barang yang paling mahal terlebih dahulu dan akan membeli barang-barang yang sedang SALE masing-masing 1 jika uang yang dimilikinya masih cukup.
4. Untuk setiap transaksi kamu harus menyimpannya ke dalam sebuah file ( *Bonus Nilai )

Code By : Jution Candra Kirana
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	var saldoAwal = 700000
	var idMember string = "123132ssd666"
	// fmt.Println("id:", idMember)
	// fmt.Println("Uang:", saldo)
	var barangSale = map[string]int{
		"Sepatu Ceebook":   1500000,
		"Baju Zoro":        500000,
		"Baju H&N":         250000,
		"Sweater Uniklooh": 175000,
		"Casing Handphone": 50000,
	}

	keys := make([]string, 0, len(barangSale))
	values := make([]int, 0, len(barangSale))

	for k, v := range barangSale {
		keys = append(keys, k)
		values = append(values, v)
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})
	// jumlahbeli := 0 //isi nilai awal 0
	saldo := 700000
	var barangBeli []string
	for k := 0; k <= len(barangSale)-1; k++ {
		if values[k] <= saldo {
			saldo -= values[k]
			// jumlahbeli++ //akan terus bertambah 1 sampai budgetnya abis
			barangBeli = append(barangBeli, keys[k])
		}
	}
	// Output
	var data map[string]interface{}
	data = map[string]interface{}{
		"id":       idMember,
		"uang":     saldoAwal,
		"barang":   barangBeli,
		"sisaUang": saldo,
	}

	fmt.Println(data)
}
