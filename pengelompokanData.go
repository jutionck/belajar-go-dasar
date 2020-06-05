package main

import "fmt"

func main() {
	hurufArray := []string{"a", "b", "a", "u", "u", "c", "a"}
	countHuruf(hurufArray)
}

func countHuruf(arr []string) []map[string]int {
	var kelompokData = []map[string]int{}
	for _, huruf := range arr {
		//  bingung append ke slice map nya disini
		// muncul yang diharapkan
		//  map[a:2] map[b:1]
		dict[huruf] = dict[huruf] + 1
	}
	fmt.Println(dict)
}
