package main

import "fmt"

func main() {
	kataInput := "jution"
	balik := reverse(kataInput)
	fmt.Println(balik)
}
func reverse(kata string) (dibalik string) {
	for i := len(kata) - 1; i >= 0; i-- {
		dibalik += kata[i : i+1] //mulai dari huruf terakhir sampai awal
	}
	return dibalik
}
