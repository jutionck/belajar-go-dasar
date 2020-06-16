package model

import (
	f "fmt"
)

func hapusMhs() {

	if totalMhs() < 1 {
		f.Println("Error: data mahasiswa masih kosong")
	} else {
		i := totalMhs()
		mhs = mhs[:i-1]
		f.Println()
		f.Println("Success: data mahasiswa yang terakhir masuk berhasil dihapus")
		f.Println(mhs)
		f.Println()
		MainMenu()
	}

}
