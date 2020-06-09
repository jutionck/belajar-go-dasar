package main

import "fmt"

func bilanganPrima(firtsNumber, lastNumber int) (cetakPrima int) {
	for bil := firtsNumber; bil <= lastNumber; bil++ {
		awalPrima := 0
		for bag := 1; bag <= lastNumber; bag++ {
			if bil%bag == 0 {
				awalPrima++
			}
		}
		if awalPrima == 2 && bil != 1 {
			fmt.Println(bil)
		}
	}
	return
}

func main() {
	var awal, akhir int
	fmt.Scan(&awal, &akhir)
	bilanganPrima(awal, akhir)
}
