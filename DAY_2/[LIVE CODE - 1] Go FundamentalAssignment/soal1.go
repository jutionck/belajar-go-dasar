package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var angka string
	// var serangan string
	fmt.Print("Masukkan angka :")
	fmt.Scan(&angka)
	splitangka := strings.Split(angka, "")
	awal, _ := strconv.Atoi(splitangka[0])
	akhir, _ := strconv.Atoi(splitangka[1])

	if awal%2 == 0 && akhir%2 != 0 {
		fmt.Println("One Zombie Down!")
	} else if awal%2 == 0 && akhir%2 == 0 {
		fmt.Println("Try Again To Attack")
	} else if awal%2 != 0 && akhir%2 != 0 {
		fmt.Println("Try Again To Attack")
	} else if awal%2 != 0 && akhir%2 == 0 {
		fmt.Println("You Are Dead")
	}

}
