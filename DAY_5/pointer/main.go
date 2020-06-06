package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 5
	scale(&a, 2)
	sqr(&a)
	fmt.Println(a)
}

func scale(angka *float64, s float64) {
	*angka = *angka * s
}

func sqr(angka *float64) {
	*angka = math.Sqrt(*angka * *angka)
}
