package model

import (
	f "fmt"
	"os"
)

// MainMenu func
func MainMenu() {
	var menu string
	f.Println("**************************************")
	f.Println("Simple APP Mahasiswa")
	f.Println("**************************************")
	f.Println("1. Add Mahasiswa")
	f.Println("2. Delete Mahasiswa")
	f.Println("3. View Mahasiswa")
	f.Println("4. Update Mahasiswa")
	f.Println("5. Exit")
	f.Println("**************************************")

	f.Printf("Masukan menu yang dipilih: ")
	f.Scan(&menu)
	switch menu {
	case "1":
		tambahMhs()
	case "2":
		hapusMhs()
	case "3":
		lihatMhs()
	case "4":
		updateMhs()
	case "5":
		f.Println("Anda berhasil keluar dari aplikasi")
		os.Exit(1)
	}
}
