package main

import "fmt"

func main() {
	var operator string
	var angkapertama, angkakedua int
	angkapertama = 30
	angkakedua = 6
	operator = "/"
	fmt.Println(HitungCalculator(angkapertama, angkakedua, operator))
}

func HitungCalculator(num1, num2 int, op string) int {
	var hasil int
	switch op {
	case "+":
		hasil = num1 + num2
	case "-":
		hasil = num1 - num2
	case "/":
		hasil = num1 / num2
	case "*":
		hasil = num1 * num2
	default:
		fmt.Println("Error!")
	}
	return hasil
}
