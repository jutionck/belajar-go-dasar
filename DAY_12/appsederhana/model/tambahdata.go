package model

import (
	"fmt"
	f "fmt"
	"os"
)

func tambahMhs() {
	if totalMhs() < 5 {
		f.Println("--------------------------------------")
		f.Println("Add Mahasiswa")
		f.Println("--------------------------------------")

		for {
			f.Print("Nama (3-20 karakter) : ")
			f.Scan(&nama)

			if len(nama) > 2 && len(nama) < 21 {
				break
			}
			f.Println("Error: nama harus 3-20 karakter")
		}

		for {
			f.Printf("Umur (min 17 tahun)  : ")
			f.Scan(&umur)
			if umur > 16 {
				break
			}
			f.Println("Error: umur minimal 17 tahun")
		}

		for {
			f.Printf("Jurusan (maks 10 karakter) : ")
			f.Scan(&jurusan)

			if len(jurusan) > 1 {
				break
			}
			f.Println("Error: jurusan maksimal 10 karakter")
		}
		simpanMhs(nama, umur, jurusan)
		return
	}

	f.Println("Error: data mahasiswa sudah penuh")
	return

}

func simpanMhs(nama string, umur int, jurusan string) {
	if totalMhs() < 5 {
		var file, _ = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
		defer file.Close()
		mhs = append(mhs, dataMhs{namaMhs: nama, umurMhs: umur, jurusanMhs: jurusan})
		s := fmt.Sprintf("%v\n", mhs)
		nambahTulisan(path, s)
		fmt.Println()
		fmt.Println("Success: data mahasiswa berhasil ditambah")
		fmt.Println()
		MainMenu()
	}
}
