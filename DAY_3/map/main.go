package main

import "fmt"

func main() {

	// var nik map[int]int
	karyawan := make(map[int]string)
	var carinik int
	// isi map
	karyawan[001] = "Jution"
	karyawan[002] = "Candra"
	// cek semua karyawan
	fmt.Println("NIK \t Nama")
	fmt.Println("--------------------")
	for nik, nama := range karyawan {
		fmt.Println(nik, "\t", nama)
	}

	// cari data
	fmt.Println("Cari dengan NIK :")
	fmt.Scan(&carinik)
	var nama, isExist = karyawan[carinik]
	if isExist {
		fmt.Println(nama)
	} else {
		fmt.Println("Not Found!")
	}
}
