package main

import "fmt"

func main() {
	kata := "satu dua dan tiga"
	printCount := countVocal(kata)
	fmt.Println("Banyaknya huruf vokal: ", printCount)
}

func countVocal(kata string) (hitung int) {
	for i := 0; i < len(kata); i++ {
		if kata[i:i+1] == "a" ||
			kata[i:i+1] == "i" ||
			kata[i:i+1] == "u" ||
			kata[i:i+1] == "e" ||
			kata[i:i+1] == "o" {
			hitung++
		}
	}
	return
}
