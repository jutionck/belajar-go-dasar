package model

import (
	f "fmt"
)

func deleteContact() {
	if totalKontak() < 1 {
		f.Println("Error: data kontak masih kosong")
	} else {
		i := totalKontak()
		contacts = contacts[:i-1]
		f.Println()
		f.Println("Success: data kontak yang terakhir masuk berhasil dihapus")
		f.Println(contacts)
		f.Println()
		MainMenu()
	}
}
