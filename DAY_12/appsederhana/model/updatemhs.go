package model

import (
	f "fmt"
	sc "strconv"
)

func updateMhs() {
	for i := 0; i < totalMhs(); i++ {
		f.Println()
		f.Println(i)
		f.Println("Nama:", mhs[i].namaMhs)
		f.Println("Umur:", mhs[i].umurMhs)
		f.Println("Jurusan:", mhs[i].jurusanMhs)
		f.Println()
	}

	var index string
	for {
		f.Printf("Masukkan index yang ingin diedit: ")
		f.Scan(&index)

		index, _ := sc.Atoi(index)
		if index < totalMhs() {
			f.Println()
			f.Println(index)
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
			prosesUpdate(index, nama, umur, jurusan)
		} else {
			f.Println("Error: data mahasiswa tidak ada")
		}
		f.Println("Error: anda memasukkan index yang salah")
	}
}

func prosesUpdate(inputIndex int, nama string, umur int, jurusan string) {
	for i := range mhs {
		if i == inputIndex {
			mhs[i].namaMhs = nama
			mhs[i].jurusanMhs = jurusan
			mhs[i].umurMhs = umur
		}
	}
	f.Println()
	f.Println("Success: data mahasiswa berhasil dirubah")
	f.Println(mhs)
	f.Println()
	MainMenu()
}
