package main

import (
	"fmt"
)

func main() {
	var awal, akhir int
	fmt.Scan(&awal, &akhir)
	isideretangka := deretangka(awal, akhir)
	// sama aja kayak
	// isiratarata, valid := rataRata(isideretangka)
	if isiratarata, valid := rataRata(isideretangka); valid {
		fmt.Println("Deret \t\t:", isideretangka)
		fmt.Println("Rata-rata \t:", isiratarata)
	} else {
		fmt.Println("Rata-rata kosong")
	}
}

func deretangka(awal, akhir int) []int {
	var hasil []int
	var err error
	if err == nil {
		for i := awal; i <= akhir; i++ {
			hasil = append(hasil, i)
		}
	} else {
		fmt.Println(err.Error())
	}
	return hasil
}

func rataRata(hasilderet []int) (float32, bool) {
	total := 0
	panjangderet := len(hasilderet)
	if panjangderet == 0 {
		return float32(total), false
	}
	for _, value := range hasilderet {
		total += value
	}
	hasilratarata := float32(total / len(hasilderet))
	return hasilratarata, true
}
