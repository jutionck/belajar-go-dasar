package model

import (
	"fmt"
	f "fmt"
	"log"
	"os"
	"regexp"
)

func addGroup() {
	// validasi grup masksimal 3
	if totalGroup() < 3 {
		f.Println("--------------------------------------")
		f.Println("            Add Group                 ")
		f.Println("--------------------------------------")

		f.Print("Group ID : ")
		f.Scan(&groupID)

		for {
			f.Print("Group Name : ")
			f.Scan(&groupName)
			var regex, _ = regexp.Compile(`[a-z]+`)
			var isMatch = regex.MatchString(groupName)
			if isMatch == true {
				break
			}
			f.Println("Error: Nama grup tidak boleh angka")
			if len(groupName) > 5 {
				break
			}
			f.Println("Error: Nama harus minimal 5 karakter")
		}

		simpanGroup(groupID, groupName)
		return
	}
}

func simpanGroup(groupID int, groupName string) {
	var file, _ = os.OpenFile(pathGroup, os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()
	groups = append(groups, group{
		id:   groupID,
		name: groupName,
	})
	s := f.Sprintf("%v\n", groups)
	nambahTulisanGroup(pathGroup, s)
	f.Println()
	f.Println("Success: data group berhasil dibuat")
	f.Println()
	MainMenu()
}

func nambahTulisanGroup(fileName string, tulisan string) {
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

	fmt.Println("==> file group berhasil terisi")

}
