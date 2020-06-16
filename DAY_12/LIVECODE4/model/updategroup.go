package model

import (
	f "fmt"
	"regexp"
	sc "strconv"
)

func updateGroup() {
	for i := 0; i < totalGroup(); i++ {
		f.Println()
		f.Println(i)
		f.Println("Nama:", groups[i].name)
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

			prosesUpdateGroup(index, groupName)
			return
		} else {
			f.Println("Error: data mahasiswa tidak ada")
		}
		f.Println("Error: anda memasukkan index yang salah")
	}
}

func prosesUpdateGroup(inputIndex int, groupName string) {
	for i := range groups {
		if i == inputIndex {
			groups[i].name = groupName
		}
	}
	f.Println()
	f.Println("Success: data group berhasil dirubah")
	f.Println(groups)
	f.Println()
	MainMenu()
}
