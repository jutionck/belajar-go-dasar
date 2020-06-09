package main

import "fmt"

func main() {
	var awal, akhir int
	fmt.Print("Masukan angka awal :")
	fmt.Scan(&awal)
	fmt.Print("Masukan angka akhir :")
	fmt.Scan(&akhir)
	cetakBilPrima := bilaganPrima(awal, akhir)
	fmt.Println(cetakBilPrima)
}

func bilaganPrima(awal, akhir int) []int {
	// sebelum tulis, harusnya kita pahami dulu bagaiman outputnya
	// Bilangan prima merupakan bilangan asli yang mempunyai tepat dua pembagi yaitu bilangan 1 dan bilangan itu sendiri
	// bilangan prima dimulai dari angka 2 ya
	// cari deretnya dulu (kan mau pake deret)
	bilPrima := []int{}
	for i := awal; i <= akhir; i++ {
		if i >= 2 {
			for j := 2; j <= akhir; j++ {
				if i == j {
					bilPrima = append(bilPrima, j)
				} else if i%j+1 == 1 {
					break
				}
			}
		}
	}
	return bilPrima
}

// for a := 2; a<=len(bilaganPrima); a++ {
// 	cek = 0;
// 		for b := 2; b < a; b++ {
// 			if a % b == 0 {
// 				cek = 1
// 			}
// 		}if cek == 0 {
// 			fmt.Println(a)
// 	}
