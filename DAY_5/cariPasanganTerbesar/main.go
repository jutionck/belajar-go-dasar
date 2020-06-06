package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	angka := 18392
	cetakHasil := pasanganAngka(angka)
	fmt.Println(cetakHasil)
}

func pasanganAngka(number int) (max int) {
	// pertama kita convert dulu si number ini ke string, karena mau kita loop (loop tidak menerima data type number)
	convertNumber := strconv.Itoa(number)
	// jika sudah di convert baru deh kita split, tujuannya adalah kita mau buat ke slice
	splitConvertNumber := strings.Split(convertNumber, "")
	// jika udah di split, kan jadi ni slice of string nya, baru kita lakukan perulangan
	// jangan lupa kita definisikan nilai_sementara, inget ini number ya
	nilaiSementara := 0
	for i := 0; i < len(splitConvertNumber)-1; i++ {
		// udah itu kita masukin hasil looping ini (string) ke nilai_sementara (tapi) jangan lupa kita convert dulu ya hasil looping ke int (kan) tadi masih slice of string hhe
		// oh ya kalo mau convert dari string ke int itu ada 2 return nya, 1 buat hasilnya, ke 2 biasanya bool, kalo gak mau di pake kita kasih ignore aja ya (_)
		nilaiSementara, _ = strconv.Atoi(splitConvertNumber[i] + splitConvertNumber[i+1])
		// buat kondisi ya
		if nilaiSementara > max {
			max = nilaiSementara
		}
	}
	// kita kembalikan nilai nya ke max, karena max udah di definisikan di func jadi cukup
	return
}
