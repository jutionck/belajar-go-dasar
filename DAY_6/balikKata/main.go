package main

import (
	"fmt"
	"strings"
)

func main() {
	var tampung string
	kataYangDibalik := "JUTION"
	// tampung = kataYangDibalik
	balikKata(&kataYangDibalik, &tampung)
	fmt.Println(tampung)
}

func balikKata(kata *string, tampung *string) {
	splitKata := strings.Split(*kata, "")
	// fmt.Println(splitKata[0:2])
	// i -= 2 (buat langkah 2)
	for i := len(splitKata) - 1; i >= 0; i-- {
		if i != 1 && len(splitKata)-1 == i {
			*tampung += splitKata[i]
		}

	}
}
