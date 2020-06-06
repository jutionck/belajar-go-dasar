/*
tolong buat menu baru untuk :
kategori produk
 1. Buat kategori baru
2. Daftar kategori
*/
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
)

var productList []map[string]string
var categoryList []map[string]string
var path = "/Users/jutioncandrakirana/Documents/GitHub/enigmaVica/DAY_5/temp/test.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func createFile() {
	// cek ada gak file nya
	var _, err = os.Stat(path)
	// kalo belum di buat
	if os.IsNotExist(err) {
		var _, err = os.Create(path)
		if isError(err) {
			return
		}
		fmt.Println("==> file berhasil dibuat", path)
	}
}

func readFile() {
	// Open file for reading.
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		// Break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// Break if error occured
		if err != nil && err != io.EOF {
			isError(err)
			break
		}
	}

	fmt.Println("Reading from file.")
	fmt.Println(string(text))
}

func writeFile(teks string) {
	// buka file dengan level akses READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// tulis data ke file
	_, err = file.WriteString(teks)
	if isError(err) {
		return
	}

	// simpan perubahan
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("==> file berhasil di isi")
}

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
func mainMenu() {
	fmt.Println("**************************************")
	fmt.Println("Aplikasi")
	fmt.Println("**************************************")
	fmt.Println("1. Buat Produk Baru")
	fmt.Println("2. Daftar Produk")
	fmt.Println("3. Buat Kategori Baru")
	fmt.Println("4. Daftar Kategori")
	fmt.Println("5. Keluar")
	fmt.Println("Pilihan menu dari 1-5")
}

func listProductForm() {
	fmt.Println("Halaman Daftar Produk")
	fmt.Println("**************************************")
	fmt.Printf("%3s \t%-20s\t%-20s\n", "No.", "Kode Produk", "Nama Produk")
	for idx, product := range productList {
		fmt.Printf("%-3d\t%-20s\t%-20s\n", idx, product["productCode"], product["productName"])
	}
}
func newProductForm() {
	var productCode string
	var productName string
	var saveProductConfirmation string
	fmt.Println("Halaman Buat Produk Baru")
	fmt.Println("**************************************")
	fmt.Print("Product Code : ")
	fmt.Scanln(&productCode)
	fmt.Print("Product Name : ")
	fmt.Scanln(&productName)
	fmt.Printf("Produk %s dengan kode %s akan disimpan (y/n) ? ", productName, productCode)
	fmt.Scanln(&saveProductConfirmation)
	if saveProductConfirmation == "y" {
		newProduct := make(map[string]string)
		newProduct["productCode"] = productCode
		newProduct["productName"] = productName
		productList = append(productList, newProduct)
		clearScreen()
	}
}
func listCategoryForm() {
	fmt.Println("Halaman Daftar Kategori")
	fmt.Println("**************************************")
	fmt.Printf("%3s \t%-20s\t%-20s\n", "No.", "Kode Kategori", "Nama Kategori")
	for idx, category := range categoryList {
		fmt.Printf("%-3d\t%-20s\t%-20s\n", idx, category["categoryCode"], category["categoryName"])
	}
}
func newCategoryForm() string {
	var categoryCode string
	var categoryName string
	var saveCategoryConfirmation string
	fmt.Println("Halaman Buat Kategori Baru")
	fmt.Println("**************************************")
	fmt.Print("Category Code : ")
	fmt.Scanln(&categoryCode)
	fmt.Print("Category Name : ")
	fmt.Scanln(&categoryName)
	fmt.Printf("kategori %s dengan kode %s akan disimpan (y/n) ? ", categoryName, categoryCode)
	fmt.Scanln(&saveCategoryConfirmation)
	if saveCategoryConfirmation == "y" {
		newCategory := make(map[string]string)
		newCategory["categoryCode"] = categoryCode
		newCategory["categoryName"] = categoryName
		categoryList = append(categoryList, newCategory)
		clearScreen()
	}
	masukanisi := new(bytes.Buffer)
	writeFile(masukanisi)
}

func main() {
	mainMenu()
	for {
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			newProductForm()
			break
		case "2":
			listProductForm()
			break
		case "3":
			newCategoryForm()
			break
		case "4":
			listCategoryForm()
			break
		case "5":
			os.Exit(0)
		default:
			clearScreen()
		}
	}
}
