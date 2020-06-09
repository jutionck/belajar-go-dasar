/*
Output = Memecahkan sebuah pola perkalian (array)
Proses = Buat function menerima string
Input = 42#3 * 188 = 80#204, hasilnya [8,5] karena 4283 * 188 = 805204
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	isiPola := "42#3 * 188 = 80#204"
	// var angkaGanti = []int{0, 1, 2, 3, 4, 5, 6, 7, 8} //di persiapkan untuk meganti karakter yang ada di pola
	polaPerkalian(isiPola)
	// fmt.Println(tempHasil)

}

func polaPerkalian(pola string) []string {
	var hasilCari = []string{}
	// find := "#"
	// replaceWith := "9"
	// var replacePola = strings.Replace(pola, find, replaceWith, 2) //mengganti karakter -> # keganti
	// fmt.Println(replacePola)
	// fmt.Println(strings.Index(pola, "#"))
	splitPolaPertama := strings.Split(pola, " ")
	ubahPertama := splitPolaPertama[0]
	// fmt.Println(ubahPertama)
	// ubahKedua := splitPolaPertama[4]
	// fmt.Println(ubahKedua)
	// pengKalinya := splitPolaPertama[2]
	// fmt.Println(pengKalinya)
	// caariPagerPertama := strings.Index(ubahPertama, "#")
	// fmt.Println(caariPagerPertama)
	// fmt.Println(strings.Index(ubahPertama, "#")) //2 true
	// caariPagerKedua := strings.Index(ubahKedua, "#")
	// fmt.Println(caariPagerKedua)
	pisahUbahPertama := strings.Split(ubahPertama, "")
	for _, value := range pisahUbahPertama {
		fmt.Print(value)
	}

	return hasilCari
}
