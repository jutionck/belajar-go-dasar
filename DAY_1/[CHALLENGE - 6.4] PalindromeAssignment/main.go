// Palindrome
package main

import (
	"fmt"
	"strings" //package untuk manipulasi string
)

func main() {
	var kata string = "jaja"
	var KataDibalik string = ""
	var JumlahKata = len(kata)
	// input
	// fmt.Print("Masukkan kata : ")
	// fmt.Scan(&kata)
	for i := JumlahKata - 1; i >= 0; i-- {
		KataDibalik = KataDibalik + string(kata[i])
	}

	// validasi apakah huruf besar kecil di samakan
	if strings.ToLower(kata) == strings.ToLower(KataDibalik) {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")
	}
}
