package model

import (
	"fmt"
	f "fmt"
	"log"
	"os"
	"regexp"
)

func addContact() {
	if totalKontak() < 5 && totalGroup() > 0 {
		f.Println("--------------------------------------")
		f.Println("              Add Contact             ")
		f.Println("--------------------------------------")

		f.Print("Contact ID : ")
		f.Scan(&contactID)

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

		simpanContact(contactID, contactName, contactGroupID)
		return

	} else {
		f.Println("Error: To add contact must be have a group")
		MainMenu()
	}
}

func simpanContact(ncontactID int, contactName string, contactGroupID int) {
	var file, _ = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()
	contacts = append(contacts, contact{
		id:      contactID,
		name:    contactName,
		groupID: contactGroupID,
	})
	s := f.Sprintf("%v\n", contacts)
	nambahTulisan(path, s)
	f.Println()
	f.Println("Success: data contact berhasil ditambah")
	f.Println()
	MainMenu()
}

func nambahTulisan(fileName string, tulisan string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer f.Close()

	if _, err = f.WriteString(tulisan); err != nil {
		log.Fatalf(err.Error())
	}

	// simpan perubahan
	err = f.Sync()
	if err != nil {
		return
	}

	fmt.Println("==> file berhasil di isi")

}
