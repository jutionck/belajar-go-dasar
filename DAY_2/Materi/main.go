package main

import (
	"fmt"
)

func main() {
	a := []string{"1", "2"}                         //slice
	fmt.Printf("Isi A : %v\n", a)                   // cetak isi slice a
	fmt.Println("len(a)", len(a), "cap(a)", cap(a)) //cek panjang dan kapasistas
	b := append(a, "3")                             //tambah baru di ambil dari a
	fmt.Printf("Isi B : %v\n", b)                   // cetak isi
	fmt.Println("len(b)", len(b), "cap(b)", cap(b)) //cek panjang dan kapasitas
	c := append(b, "4")                             // tambah value yang di ambil dari b
	fmt.Printf("Isi C : %v\n", c)                   //cetak isi value
	fmt.Println("len(c)", len(c), "cap(c)", cap(c)) //cek panjang dan kapasitas
	d := append(b, "5")
	fmt.Printf("Isi D : %v\n", d)
	fmt.Println("len(d)", len(d), "cap(d)", cap(d))
	fmt.Println("Lihatlah isi C menjadi")
	fmt.Printf("Isi C : %v\n", c)
}
