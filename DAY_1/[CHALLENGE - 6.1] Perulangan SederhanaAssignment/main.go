// Perulangan Sederhana
package main

import (
	"fmt"
)

func main() {
	var count int
	fmt.Printf("Masukan Jumlah Perulangan :")
	fmt.Scan(&count)
	for i := count; i > 0; i-- {
		fmt.Println(i, "- I will become a go developer")
	}

}
