/*
1. Isi coin adalah 1,5,7,9,11 masing-masing koin di batasi 1
2. Function menerima inputan number, ex : 12, 17 dll
3. Mengembalikan nilai yang jumlahnya sama kyk yang di inputkan (unik tidak boleh sama)
4. Hasil, misal input 12 -> keluarnya [1,11] [5,7]

*/
package main

import "fmt"

func main() {
	var isiCoin = []int{1, 5, 7, 9, 11}
	inputNumber := 12
	if inputNumber > 33 {
		fmt.Println("Maksimal adalah 33 tidak lebih")
	}
	pasanganCoin(isiCoin, inputNumber)
}

func pasanganCoin(isiCoin []int, sumNumber int) {
	for i := 0; i < len(isiCoin); i++ {
		for j := i; j < len(isiCoin); j++ {
			if isiCoin[i]+isiCoin[j] == sumNumber {
				fmt.Printf("[[%d %d]]", isiCoin[i], isiCoin[j])
				break
			}
			for k := j; k < len(isiCoin); k++ {
				if isiCoin[i]+isiCoin[j]+isiCoin[k] == sumNumber {
					// fmt.Println(isiCoin[i], isiCoin[j], isiCoin[k])
					fmt.Printf("[[%d %d %d]]", isiCoin[i], isiCoin[j], isiCoin[k])
					break
				}
			}
		}
	}
}
