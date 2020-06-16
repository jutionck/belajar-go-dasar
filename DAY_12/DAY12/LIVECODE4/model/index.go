package model

import (
	f "fmt"
	"os"
)

// MainMenu func
func MainMenu() {
	var menu string
	f.Println("***********************************")
	f.Println("    Main Menu [Add Contact]        ")
	f.Println("***********************************")
	f.Println("1. Add Contact")
	f.Println("2. Add Group")
	f.Println("3. View List Contact")
	f.Println("4. View List Group")
	f.Println("5. Show Contact With Group Name")
	f.Println("6. Delete Contact")
	f.Println("7. Delete Group")
	f.Println("8. Update Contact")
	f.Println("9. Update Group")
	f.Println("10. Logout")
	f.Println("***********************************")

	// validasi menu yang di pilih harus 1-5
	f.Printf("Masukan menu yang dipilih (1-10): ")
	f.Scan(&menu)
	switch menu {
	case "1":
		addContact()
	case "2":
		addGroup()
	case "3":
		viewContact()
	case "4":
		viewGroup()
	case "5":
		viewContactGroup()
	case "6":
		deleteContact()
	case "7":
		deleteGroup()
	case "8":
		updateContact()
	case "9":
		updateGroup()
	case "10":
		f.Println("Logout Success!!")
		os.Exit(1)
	default:
		f.Println("Menu salah!!")
	}
}
