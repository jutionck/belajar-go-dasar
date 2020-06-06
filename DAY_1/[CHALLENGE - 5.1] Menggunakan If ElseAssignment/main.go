// Menggunakan If Else
package main

import (
	"fmt"
)

func main() {
	var nama string = ""
	var peran string = ""

	// kode disini
	if nama == "" && peran == "" {
		fmt.Print("Nama dan Peran Harus Di Isi")
	} else if peran == "" {
		fmt.Print("Peran Harus Di Isi")
	} else if peran == "superhero" {
		fmt.Print("Selamat Datang Superhero Saitama, Kalahkan Semua Monster Di Muka Bumi")
	} else if peran == "monster" {
		fmt.Print("Selamat Datang Monster Saitama, Hancurkan Semua Superhero Yang Ada")
	} else {
		fmt.Print("Selamat Datang Saitama, Pilih Peranmu Untuk Melanjutkan Game Ini")
	}
}
