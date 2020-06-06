package main

import "fmt"

func main() {
	panjang, lebar := 10, 10
	luas := 0
	luasBalok(&luas, &panjang, &lebar)
	fmt.Println(&luas)
}
func luasBalok(luas, panjang, lebar *int) {
	*luas = *panjang * *lebar
}
