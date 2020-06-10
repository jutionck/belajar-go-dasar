package main

import "fmt"

func main() {
	var beratBadan kilogram
	beratBadan = kilogram(10)
	fmt.Println(beratBadan)

	var beratShabu gram
	beratShabu = gram(1)
	fmt.Println(beratShabu)

	total := beratBadan + gramToKilogram(beratShabu)
	fmt.Println(total)
}
