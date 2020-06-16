// Buatlah sebuah aplikasi sederhana tentang cara memesan makanan di sebuah restoran. Bisa kalian bayangkan jika kalian ingin memesan makanan tetapi harus menunggu makanan pelanggan lain selesai terlebih dahulu baru kalian bisa memesan makanan. Tugas kalian adalah membuat setiap pesanan pelanggan menjadi tidak saling tunggu.
// Misalkan jika budi memesan makanan pertama dengan menu nasi goreng, ayam goreng dan jus alpukat . Lalu setelahnya tono hanya memesan makanan nasi goreng saja . Maka menurut kalian pesanan siapa dulu yang akan selesai ? pastinya pesanan tono yang akan selesai karna tono hanya memesan nasi goreng walaupun tono memesan setelah budi .
// List Makanan :
// 1. Nasi Goreng Dengan Waktu Pembuatan 4 Detik
// 2. Spaghetti Dengan Waktu Pembuatan 3 Detik
// 3. Ayam Goreng Dengan Waktu Pembuatan 5 Detik
// 4. Jus Alpukat Dengan Waktu Pembuatan 2 Detik
// 5. Air Mineral Dengan Waktu Pembuatan 1 Detik

package main

import (
	"fmt"
	"time"
)

type makanan struct {
	namaMakanan    string
	waktuPenyajian int
}
type pembeli struct {
	namaPembeli string
	daftarBeli  []makanan
}
type pesanan struct {
	namaYangPembeli string
	lamaMemasak     int
}

var daftarMenu = []makanan{
	{"nasi goreng", 4},
	{"spaghetti", 3},
	{"ayam goreng", 5},
	{"Jus alpukat", 2},
	{"air mineral", 1},
}
var daftarPesanan []pesanan

func main() {
	c := make(chan string)
	var waktuPenyajian int
	var budi = pembeli{"Budi", []makanan{{"nasi goreng", 1}, {"spaghetti", 2}}}
	var tono = pembeli{"Tono", []makanan{{"nasi goreng", 1}}}
	var daftarPembeli = []pembeli{budi, tono}
	for _, isi := range daftarPembeli {
		for _, beli := range isi.daftarBeli {
			for _, daftar := range daftarMenu {
				if daftar.namaMakanan == beli.namaMakanan {
					temp := beli.waktuPenyajian * daftar.waktuPenyajian
					waktuPenyajian += temp
					temp = 0
				}
			}
		}
		pemesan := pesanan{
			namaYangPembeli: isi.namaPembeli,
			lamaMemasak:     waktuPenyajian,
		}
		daftarPesanan = append(daftarPesanan, pemesan)
		waktuPenyajian = 0
	}
	for _, isi := range daftarPesanan {
		go masakMakanan(isi.namaYangPembeli, isi.lamaMemasak, c)
	}
	for i := 0; i < len(daftarPesanan); i++ {
		fmt.Println(<-c)
	}
}

func masakMakanan(inNamaPembeli string, inWaktuPenyajian int, c chan string) {
	time.Sleep(time.Second * time.Duration(inWaktuPenyajian))
	fmt.Println("Selamat Menikmati ", inNamaPembeli)
	c <- "jalan"
}
