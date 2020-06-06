package main

import ("fmt")

func main() {
	var input, ganjil, genap
	fmt.Print("Masukkan angka :")
	fmt.Scan(&input)
	fmt.Print("Ganjil : ")
	for i :=1; i < input; i++ {
		if i%2 != 0 {
			if i  != 1 {
				fmt.Print(" + ")
			}
			fmt.Print(i)
			ganjil += i
		}
	}
	fmt.Print(" = ", ganjil)
}