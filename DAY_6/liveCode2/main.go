package main

import (
	"fmt"
	"strings"
)

func main() {
	input1 := "42#3"
	input2 := "188"
	hasil := "80#204"
	hasilCari := cariPagar(input1, input2, hasil)
	fmt.Println(hasilCari)
}
func cariPagar(input1, input2, hasil string) []string {
	sliceStr1 := strings.Split(input1, "")
	// sliceStr2 := strings.Split(input2, "")
	sliceHasil := strings.Split(hasil, "")
	inputBenar := "4283"
	hasilBenar := "805204"
	sliceInptBnr := strings.Split(inputBenar, "")
	sliceHslBnr := strings.Split(hasilBenar, "")
	var hasilCari = []string{}
	for index, isi := range sliceStr1 {
		if isi == "#" {
			for i := 0; i < len(sliceInptBnr); i++ {
				// fmt.Println(isi, sliceInptBnr[i])
				if index == i {
					// fmt.Println(sliceInptBnr[i])
					hasilCari = append(hasilCari, sliceInptBnr[i])
					// break
					// fmt.Println(hasilCari)
				}

			}
		}

	}
	for index1, isiHasil := range sliceHasil {
		if isiHasil == "#" {
			for j := 0; j < len(sliceHslBnr); j++ {
				// fmt.Println(isi, sliceInptBnr[i])
				if index1 == j {
					// fmt.Println(sliceInptBnr[i])
					hasilCari = append(hasilCari, sliceHslBnr[j])
					// break
					// fmt.Println(hasilCari)
				}

			}

		}
	}
	return hasilCari

}
