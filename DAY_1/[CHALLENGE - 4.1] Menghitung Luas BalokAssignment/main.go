// Menghitung Luas Balok
package main

import (
	"fmt"
)

func main() {
	var (
		panjang int
		lebar   int
		tinggi  int
		LBalok  int
		VBalok  int
	)
	fmt.Printf("Masukkan panjang : ")
	fmt.Scan(&panjang)
	fmt.Printf("Masukkan lebar : ")
	fmt.Scan(&lebar)
	fmt.Printf("Masukkan tinggi : ")
	fmt.Scan(&tinggi)
	LBalok = panjang * lebar
	VBalok = panjang * lebar * tinggi
	fmt.Println("Luas permukaan : ", LBalok)
	fmt.Println("Luas Volume : ", VBalok)
}
