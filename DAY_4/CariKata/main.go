package main

import "fmt"

func main() {
	isi := []string{"a", "3", "c"}
	cari := "3"
	tampungPencarian, index := cariKata(isi, cari)
	fmt.Println(tampungPencarian, index)
}

func cariKata(isi []string, cari string) (bool, int) {
	var cek bool
	var index int
	for i := 0; i < len(isi); i++ {
		if isi[i] == cari {
			cek = true
			index = i
		} else {
			cek = false
			index = -1
		}
	}
	return cek, index
}
