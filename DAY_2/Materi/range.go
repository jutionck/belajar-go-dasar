package main

import "fmt"

func main() {
	var namadicari string
	var kumpulannama = []string{"Budi", "Tono", "Sasa", "Mahmud"}
	fmt.Print("Masukkan nama yang akan di tampilkan :")
	fmt.Scan(&namadicari)
	// fmt.Println(kumpulannama)
	for _, nama := range kumpulannama {
		if namadicari == nama {
			fmt.Println(namadicari)
			break
		} else {
			fmt.Println("Data tidak ada")
			break
		}
	}
}
