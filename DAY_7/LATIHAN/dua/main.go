/*
Method adalah fungsi yang menempel pada type (struct bisa yang lain)
Uji coba method, script by Kang Prana
*/

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// type student struct { //struct student yang menempel di func fullname
// 	firtsName, lastName string
// 	age                 int
// }

// func (s student) fullname() string {
// 	return s.firtsName + " " + s.lastName
// }

// func (s student) getNameAt(age int) string {
// 	gabungNama := s.firtsName + " " + s.lastName
// 	return strings.Split(gabungNama, " ")[age-1]
// }

// func main() {
// 	student1 := student{
// 		"Budi",
// 		"Anduk",
// 		18,
// 	}
// 	namaLengkap := student1.fullname()
// 	fmt.Println(namaLengkap)
// 	name := student1.getNameAt(3)
// 	fmt.Println(name)
// }

/*
Contoh lain dari penulisan method
by : Bang Edward
*/
// package main

// import (
// 	"fmt"
// 	"strings"
// 	)

// type kata string

// func (k kata) balikKata() string {
// 	var kataTerpisah = strings.Split(string(k), "")
// 	var hasil []string
// 	for i := len(kataTerpisah) - 1; i >= 0; i-- {
// 		hasil = append(hasil, kataTerpisah[i])
// 		}
// 		return strings.Join(hasil, "")
// }
// func (k kata) panjangKata() int {
// 	return len(k)
// }

// func main() {
// 	var kataSaya kata = kata("Edo")
// 	fmt.Println(kataSaya.balikKata())
// 	fmt.Println(kataSaya.panjangKata())
// }

//contoh 2
package main

import "fmt"

type bujurSangkar struct { //struct bujurSangkar isi nya sisi
	sisi      float64
	luasBujur float64
	kelBujur  float64
}

// ambil struct yang dibuat, dengan nama method nya luas() dengan nilai kembalinya adalah float64
//fungsi receiver (pointer)
func (b *bujurSangkar) luas() {
	b.luasBujur = b.sisi * b.sisi
}

func (b *bujurSangkar) keliling() {
	b.kelBujur = 4 * b.sisi
}

func main() {
	//buat variabel baru buat nampung si struct yang di buat dengan isian sisi di isi 10.0
	// var bujurSaya = bujurSangkar{sisi: 10.0}
	// cetak hasilnya dengan memanggil methodnya
	// fmt.Println(bujurSaya.luas())
	// pake pointer
	var bujurSaya = bujurSangkar{sisi: 10.0}
	bujurSaya.luas()
	bujurSaya.keliling()
	fmt.Println(bujurSaya.luasBujur) //isi pointernya
	fmt.Println(bujurSaya.kelBujur)  //isi pointernya
}
