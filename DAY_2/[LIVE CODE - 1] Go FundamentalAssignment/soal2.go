package main

import (
	"fmt"
	"strconv"
)

func main() {
	angka := "-1-1-1-1-1-2"
	value, _ := strconv.Atoi(angka)
	fmt.Println(angka)
	fmt.Println(value)
}
