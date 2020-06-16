// Program mencetak deret bilanga Genap, Ganjil, dan Kelipatan 3
// Oleh : Jution Candra Kirana

package main

import "fmt"

func cetakDeret(angka []int) [][][]int {
	var genap []int
	for _, val := range angka {
		if val%2 == 0 && val%3 != 0 {
			genap = append(genap, val)
		}
	}
	var ganjil []int
	for _, val := range angka {
		if val%2 != 0 && val%3 != 0 {
			ganjil = append(ganjil, val)
		}
	}
	var kelTiga []int
	for _, val := range angka {
		if val%3 == 0 {
			kelTiga = append(kelTiga, val)
		}
	}
	var hasil = [][][]int{{genap}, {ganjil}, {kelTiga}}
	return hasil
}

func main() {
	var isiDeret = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	tempHasil := cetakDeret(isiDeret)
	fmt.Println(tempHasil)
}
