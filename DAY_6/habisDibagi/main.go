/*
1. Jika angka tidak habis dibagi 2 atau 3, maka function akan me-return "SALAH".
2. Jika angka habis dibagi 2, maka return "TES"
3. Jika angka habis dibagi 3, maka return "LA"
4. Jika angka habis dibagi 2 dan 3, maka return "TESLA"
*/

package main

import "fmt"

func main() {
	fmt.Println(habisDibagi(10))
}

func habisDibagi(angka int) string {
	if angka%2 != 0 || angka%3 != 0 {
		return "SALAH"
	} else if angka%2 == 0 {
		return "TES"
	} else if angka%3 == 0 {
		return "LA"
	} else if angka%2 == 0 && angka%3 == 0 {
		return "TESLA"
	} else {
		return "NOT FOUND"
	}
}
