package main

import "fmt"

type mobil struct {
	merk  string
	ban   int
	spion int
	warna string
}

func main() {

	var allMobil = []mobil{
		{merk: "Toyota", ban: 4, spion: 2, warna: "Merah"},
		{merk: "Honda", ban: 8, spion: 4, warna: "Biru"},
	}

	for _, value := range allMobil {
		// fmt.Println(value)
		if value.spion > 2 && value.warna == "Biru" {
			fmt.Println(value)
		}
	}

}
