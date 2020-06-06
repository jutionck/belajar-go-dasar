// Bermain Dengan Variabel dan Tipe DataAssignment
package main

import (
	"fmt"
)

func main() {
	var a int = 15
	var b int = 20
	// kode disini
	var c int
	c = b
	b = a
	a = c
	fmt.Println("a = ", c, "b = ", b)
}
