package model

import (
	"fmt"
	f "fmt"
	"io/ioutil"
	sc "strconv"
)

func lihatMhs() {

	if totalMhs() > 0 {
		f.Println("1.View by index")
		f.Println("2.View all data")

		var menu string
		f.Printf("Masukan menu yang dipilih: ")
		f.Scan(&menu)
		switch menu {
		case "1":
			viewByIndex()
			MainMenu()
		case "2":
			viewAllMhs()
			MainMenu()
		}
	} else {
		f.Println("Pesan: belum ada data mahasiswa")
	}

}

func viewAllMhs() {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic("Failed Opened !")
	}
	str := string(b)
	fmt.Println(str)
}

func viewByIndex() {
	var index string
	for {
		f.Printf("Masukkan index yang ingin ditampilkan: ")
		f.Scan(&index)
		index, _ := sc.Atoi(index)
		if index < totalMhs() {
			f.Println()
			f.Println(index)
			f.Println("Nama:", mhs[index].namaMhs)
			f.Println("Umur:", mhs[index].umurMhs)
			f.Println("Jurusan:", mhs[index].jurusanMhs)
			f.Println()
			break
		} else {
			f.Println("Pesan: data mahasiswa tidak ada")
		}
	}
}
