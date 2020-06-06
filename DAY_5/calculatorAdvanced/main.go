package main

import "fmt"

func calculator(ops string) func(int, int) int {
	switch ops {
	case "+":
		return func(a int, b int) int { return a + b }
	case "-":
		return func(a int, b int) int { return a - b }
	default:
		return nil
	}
}
func main() {
	f := calculator(".")
	if f != nil {
		fmt.Println(f(1, 3))
	} else {
		fmt.Println("Operator Tidak dikenal")
	}
}
