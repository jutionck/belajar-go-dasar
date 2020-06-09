/*
1. Jika angka tidak habis dibagi 2 atau 3, maka function akan me-return "SALAH".
2. Jika angka habis dibagi 2, maka return "TES"
3. Jika angka habis dibagi 3, maka return "LA"
4. Jika angka habis dibagi 2 dan 3, maka return "TESLA"
*/

package main

import "fmt"

func main() {
	angka := 14
	var hasil string
	habisDibagi(&angka, &hasil)
	fmt.Println(&angka)
	fmt.Println(&hasil)
}

func habisDibagi(angka *int, hasil *string) {
	switch {
	case *angka%2 != 0 || *angka%3 != 0:
		*hasil = "SALAH"
	case *angka%2 == 0:
		*hasil = "TES"
	case *angka%3 == 0:
		*hasil = "LA"
	case *angka%2 == 0 && *angka%3 == 0:
		*hasil = "TESLA"
	default:
		*hasil = "TAK ADA YA"
	}
}
