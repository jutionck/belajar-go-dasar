package main

import "fmt"

func main() {
	var awal, akhir int
	fmt.Print("Masukan angka awal :")
	fmt.Scan(&awal)
	fmt.Print("Masukan angka akhir :")
	fmt.Scan(&akhir)
	cetakDeret := bilaganPrima(awal, akhir)
	fmt.Println(cetakDeret)
}

func bilaganPrima(awal, akhir int) []int {
	// sebelum tulis, harusnya kita pahami dulu bagaiman outputnya
	// Bilangan prima merupakan bilangan asli yang mempunyai tepat dua pembagi yaitu bilangan 1 dan bilangan itu sendiri
	// bilangan prima dimulai dari angka 2 ya
	// cari deretnya dulu (kan mau pake deret)
	var bilPrima []int
	for i := awal; i <= akhir; i++ {
		bilPrima = append(bilPrima, i)
	}
	return bilPrima
}
