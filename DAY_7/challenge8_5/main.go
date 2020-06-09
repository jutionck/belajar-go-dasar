/*
[CHALLENGE - 8.5] - Aplikasi Basis Data Sederhanan berbasis Console
Di buat oleh : Jution Candra Kirana
Problem :
1. Bisa menampung sampai dengan 5 mahasiswa
2. Validasi sesuai dengan screen, jika validasi tidak terpenuhi maka input tersebut akan terus diminta sampai terpenuhi
3. Design screen harus mempunyai menu tambah mahasiswa, hapus mahasiswa, lihat mahasiswa dan exit
4. Untuk menu tambah mahasiswa berikan validasi yaitu nama harus minimal 3 dan maksimal 20 karakter, umur minimal harus 17 tahun, jurusan maksimal 10 karakter
5. Untuk menu hapus mahasiswa akan otomatis index terakhir yang terhapus
6. untuk menu lihat mahasiswa akan ada sub menu lagi yaitu lihat menggunakan index dan lihat semua
7. untuk menu lihat menggunakan index di input sesuai index data mahasiswa
8. untuk menu lihat semua akan menampilkan seluruh data mahasiswa yang terdaftar
*/
package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type mahasiswa struct {
	namaMhs, jurusanMhs string
	umurMhs             int
}

var mahasiswaListAll = []mahasiswa{}

func clearScreen() {
	osRunning := runtime.GOOS
	if osRunning == "linux" || osRunning == "darwin" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if osRunning == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	mainMenu()
}

// main menu nya
func mainMenu() {
	fmt.Println("**************************************")
	fmt.Println("Aplikasi Basis Data Sederhana")
	fmt.Println("**************************************")
	fmt.Println("1. Add Mahasiswa")
	fmt.Println("2. Delete Mahasiswa")
	fmt.Println("3. View Mahasiswa")
	fmt.Println("4. Exit")
	fmt.Println("Pilihan menu dari 1-4")
}
func menuView() {
	fmt.Println("**************************************")
	fmt.Println("View Mahsiswa")
	fmt.Println("**************************************")
	fmt.Println("1. View By Index")
	fmt.Println("2. View All")
	fmt.Println("3. Exit")
	fmt.Println("Pilihan menu dari 1-3")
}

func lihatMahasiswa() {
	menuView()
	for {
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			mahasiswaByIndex()
			break
		case "2":
			semuaMahsiswa()
			break
		case "3":
			os.Exit(0)
		default:
			mainMenu()
		}
	}
}

func semuaMahsiswa() {
	fmt.Println("VIEW ALL")
	fmt.Println("**************************************")
	fmt.Printf("%3s \t%-20s\t%-20s\t%-20s\n", "No.", "Nama Mahasiswa", "Umur", "Jurusan")
	for idx, mhs := range mahasiswaListAll {
		fmt.Printf("%-3d\t%-20s\t%-20d\t%-20s\n", idx, mhs.namaMhs, mhs.umurMhs, mhs.jurusanMhs)
	}
	mainMenu()
}
func mahasiswaByIndex() {
	var inputIndex int
	fmt.Print("Masukkan index mahasiswa : ")
	fmt.Scanln(&inputIndex)
	for index, mhs := range mahasiswaListAll {
		if index == inputIndex-1 {
			fmt.Println("**************************************")
			fmt.Println("VIEW BY INDEX")
			fmt.Println("**************************************")
			fmt.Println("Nama : ", mhs.namaMhs)
			fmt.Println("Umur : ", mhs.umurMhs)
			fmt.Println("Jurusan : ", mhs.jurusanMhs)
			fmt.Println("--------------------------------------")
		} else {
			fmt.Println("Data tidak di temukan")
		}
	}
	mainMenu()
}

func tambahMahasiswa() {
	var namaMhs, jurusanMhs, konfirmasiMahasiswa string
	var umurMhs int
	fmt.Println("Form Add Mahasiswa")
	fmt.Println("**************************************")
	fmt.Print("Nama Mahasiswa (3-20 Karakter) \t: ")
	fmt.Scanln(&namaMhs)
	fmt.Print("Umur Mahasiswa (Minimal 17 Tahun) : ")
	fmt.Scanln(&umurMhs)
	fmt.Print("Jurusan (Maksimal 10 Karakter \t\t:")
	fmt.Scanln(&jurusanMhs)
	if len(namaMhs) >= 3 && len(namaMhs) <= 20 && umurMhs >= 17 && len(jurusanMhs) <= 10 {
		fmt.Printf("Mahasiswa %s dengan umur %d  dan jurusan %s akan disimpan (y/n) ? ", namaMhs, umurMhs, jurusanMhs)
		fmt.Scanln(&konfirmasiMahasiswa)
		if konfirmasiMahasiswa == "y" {
			newMahaiswa := mahasiswa{
				namaMhs:    namaMhs,
				umurMhs:    umurMhs,
				jurusanMhs: jurusanMhs,
			}
			mahasiswaListAll = append(mahasiswaListAll, newMahaiswa)
			clearScreen()
		} else {
			clearScreen()
		}
	} else {
		fmt.Print("Inputan tidak sesuai")
		tambahMahasiswa()
	}
}

func hapusMahasiswa() {
	fmt.Print("Hapus Mahasiswa")
	fmt.Println("**************************************")
	mahasiswaListAll[len(mahasiswaListAll)-1] = mahasiswa{"", "", 0}
	mahasiswaListAll = mahasiswaListAll[:len(mahasiswaListAll)-1]
	semuaMahsiswa() // cek hasilnya setelah di hapus
	// clearScreen()
}

func main() {
	mainMenu()
	for {
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			tambahMahasiswa()
			break
		case "2":
			hapusMahasiswa()
			break
		case "3":
			lihatMahasiswa()
			break
		case "4":
			os.Exit(0)
		default:
			clearScreen()
		}
	}
}
