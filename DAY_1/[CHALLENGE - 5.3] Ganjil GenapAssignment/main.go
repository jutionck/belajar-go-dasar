// Ganjil Genap
package main

import (
	"fmt"
)

func main() {
	var angka int
	fmt.Printf("Masukan Angka : ")
	fmt.Scan(&angka)

	// kondisi
	if angka%2 == 0 {
		fmt.Println("Genap")
	} else {
		fmt.Println("Ganjil")
	}
}
