package main

import "fmt"

func penjumalah(num1, num2 int) {
	hasil := num1 + num2
	fmt.Println(hasil)
}

func pengurangan(num1, num2 int) {
	hasil := num1 - num2
	fmt.Println(hasil)
}

func calBasic(num1, num2 int, operator string) {
	switch operator {
	case "+":
		penjumalah(num1, num2)
	case "-":
		pengurangan(num1, num2)
	default:
		fmt.Println("Belum ada operasi")
	}

}

func main() {
	var angkaPertama, angkaKedua int
	var operasi string
	fmt.Println("Angka 1: ")
	fmt.Scan(&angkaPertama)
	fmt.Println("Angka 2: ")
	fmt.Scan(&angkaKedua)
	fmt.Println("Operasi (+ - / *)")
	fmt.Scan(&angkaKedua)
	calBasic(angkaPertama, angkaKedua, operasi)
}
