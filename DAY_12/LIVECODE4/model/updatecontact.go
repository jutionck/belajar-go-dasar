package model

import (
	f "fmt"
	"regexp"
	sc "strconv"
)

func updateContact() {
	for i := 0; i < totalKontak(); i++ {
		f.Println()
		f.Println(i)
		f.Println("Nama:", contacts[i].name)
		f.Println()
	}

	var index string
	for {
		f.Printf("Masukkan index yang ingin diedit: ")
		f.Scan(&index)

		index, _ := sc.Atoi(index)
		if index < totalKontak() {
			f.Println()
			f.Println(index)
			for {
				f.Print("Nama (Min 3 karakter) : ")
				f.Scan(&contactName)
				var regex, _ = regexp.Compile(`[a-z]+`)
				var isMatch = regex.MatchString(contactName)
				if isMatch == true {
					break
				}
				f.Println("Error: Nama tidak boleh angka")
				if len(contactName) > 2 {
					break
				}
				f.Println("Error: Nama harus minimal 3 karakter")
			}

			f.Printf("Group ID : ")
			f.Scan(&contactGroupID)

			prosesUpdateContact(index, contactName, contactGroupID)
			return
		} else {
			f.Println("Error: data mahasiswa tidak ada")
		}
		f.Println("Error: anda memasukkan index yang salah")
	}
}

func prosesUpdateContact(inputIndex int, contactName string, contactGroupID int) {
	for i := range contacts {
		if i == inputIndex {
			contacts[i].name = contactName
			contacts[i].groupID = contactGroupID
		}
	}
	f.Println()
	f.Println("Success: data kontak berhasil dirubah")
	f.Println(contacts)
	f.Println()
	MainMenu()
}
