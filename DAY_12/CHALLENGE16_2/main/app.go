package main

import (
	"CHALLENGE16_2/model"
	f "fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(8)
	f.Println("RM Jution ACC")
	f.Println("List Makanan :\n1. Nasi Goreng Dengan Waktu Pembuatan 4 Detik\n2. Spaghetti Dengan Waktu Pembuatan 3 Detik\n3. Ayam Goreng Dengan Waktu Pembuatan 5 Detik\n4. Jus Alpukat Dengan Waktu Pembuatan 2 Detik\n5. Air Mineral Dengan Waktu Pembuatan 1 Detik")
	model.PrintResult()
}
