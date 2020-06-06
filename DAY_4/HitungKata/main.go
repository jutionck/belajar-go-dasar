package main

import (
	"fmt"
	"strings"
)

func main() {
	var katakata string
	fmt.Println("Masukkan kata :")
	fmt.Scan(&katakata)
	tempWordCount := wordCount(katakata)
	fmt.Println(tempWordCount)
}

func wordCount(words string) int {
	wordBreak := strings.Split(words, " ")
	hitung := len(wordBreak)
	return hitung
}
