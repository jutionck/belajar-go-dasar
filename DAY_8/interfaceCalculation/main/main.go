package main

import (
	"fmt"
	"interfaceCalculation/calculation"
)

func main() {
	var input1, input2 int
	fmt.Print("Input angka 1 : ")
	fmt.Scan(&input1)
	fmt.Print("Input angka 2 : ")
	fmt.Scan(&input2)
	substraction := calculation.Pengurangan{input1, input2}
	summary := calculation.Penjumlahan{input1, input2}
	printHasilHitung(substraction)
	printHasilHitung(summary)
}

func printHasilHitung(c calculation.Calculator) {
	hasil, pesanError := c.GetCalculation()
	if pesanError == nil {
		fmt.Println("Hasil : ", hasil)
	} else {
		fmt.Println(pesanError.Error())
	}

}
