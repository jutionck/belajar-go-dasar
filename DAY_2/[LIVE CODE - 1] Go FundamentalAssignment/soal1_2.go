package main

import "fmt"

func main() {
	var angka = [...]int{20}
	pertama := angka[0]
	if angka[0]%2 == 0 && angka[1]%2 != 0 {
		fmt.Println("One Zombie Down!")
	} else if angka[0]%2 == 0 && angka[1]%2 == 0 {
		fmt.Println("Try Again To Attack")
	} else if angka[0]%2 != 0 && angka[0]%2 != 0 {
		fmt.Println("Try Again To Attack")
	} else if angka[0]%2 != 0 && angka[0]%2 == 0 {
		fmt.Println("You Are Dead!")
	}

}
