// Palindrome next number input
// Jution Candra Kirana

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func palindromeCek(input int) {
	var (
		stringTemp, stringTemp2 string
		arrayString             []string
	)
	// Looping next number
	for i := input + 1; i <= input+input; i++ {
		temp := strconv.Itoa(i)                 //convert to string
		arrayString = append(arrayString, temp) //input to slice arrayString
		for _, isi := range arrayString {
			stringTemp = isi                           //result loop
			sliceTemp := strings.Split(stringTemp, "") //break number
			fmt.Println(sliceTemp)
			for j := len(sliceTemp) - 1; j >= 0; j-- { // polindrome check
				stringTemp2 += sliceTemp[j]
				if stringTemp == stringTemp2 {
					fmt.Println(stringTemp)
					os.Exit(0)
				}
			}
			stringTemp2 = "" //if result not print
		}
	}
}

func main() {
	number := 103
	palindromeCek(number)
}
