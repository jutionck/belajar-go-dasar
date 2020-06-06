package main

import "fmt"

func main() {
	hurufArray := []string{"a", "b", "a", "u", "u", "c", "a"}
	countHuruf(hurufArray)
}

func countHuruf(arr []string) {
	dict := make(map[string]int)
	for _, huruf := range arr {
		dict[huruf] = dict[huruf] + 1
	}
	fmt.Println(dict)
}
