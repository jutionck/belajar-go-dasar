package model

import (
	"fmt"
	f "fmt"
)

func viewContactGroup() {
	var index string
	for key, valGrup := range groups {
		fmt.Println("ID Grup : ", key)
		fmt.Println("Nama Grup : ", valGrup)
	}
	for {
		f.Printf("Masukkan id grup yang akan di cari : ")
		f.Println("List Grup : ---- ")
		f.Scan(&groupID)
		if groupID < totalKontak() {
			f.Println()
			f.Println(index)
			f.Println("Nama:", contacts[groupID].name)
			f.Println()
			break
		} else {
			f.Println("Pesan: data kontak tidak ada")
		}
	}
}
