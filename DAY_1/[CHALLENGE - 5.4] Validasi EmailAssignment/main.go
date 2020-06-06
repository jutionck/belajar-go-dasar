// Validasi Email
package main

import (
	"fmt"
	"strings"
)

func main() {
	var email string = ""
	fmt.Print("Masukkan email: ")
	fmt.Scan(&email)

	if strings.ContainsAny(email, "@") != true {
		fmt.Println("email harus mengandung @")
	} else if strings.ContainsAny(email, "@") == true {
		fmt.Println("email benar")
	} else if strings.ContainsAny(email, "@ & , ") == true {
		fmt.Println("email tidak boleh ada tanda koma")
	} else if email == "" {
		fmt.Println("email tidak boleh kosong")
	}

}
