// Penjumlahan Pasangan
package mai

import "fmt"

func main() {
	var deret1 = []int{5, 4, 1, 8, 2, 7}
	var sum int = 9
	PenjumalahPasangan(deret1, sum)
}

func PenjumalahPasangan(angka []int, total int) {
	for i := 0; i < len(angka); i++ {
		for j := i; j < len(angka); j++ {
			if angka[i]+angka[j] == total {
				fmt.Printf("[[%d %d]]", angka[i], angka[j])
			}
		}
	}
}
