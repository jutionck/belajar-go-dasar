package main

import (
	"DICTIONARYINTERFACE/pengurangan"
	"fmt"
)

func main() {
	var input1, input2 int
	fmt.Print("Input angka 1 :")
	fmt.Scan(&input1)
	fmt.Print("Input angka 2 :")
	fmt.Scan(&input2)
	substraction := pengurangan.Pengurangan{input1, input2}
	printHasilHitung(substraction)
}

func pintHasilHitung(c pengurangan.Calculator) {
	hasil, pesanMasalah
}
// func main() {
// 	// enDict := dictionary.EnglishDict{}
// 	// jpDict := dictionary.JapanDict{}

// 	// dictionaries := []dictionary.Dict{enDict, jpDict}

// 	// for _, dictionary := range dictionaries {
// 	// 	printGreeting(dictionary)
// 	// }
// 	result, err := pembagian(1, -1)
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	} else {
// 		log.Println(result)
// 	}
// }

// func pembagian(num1, num2 int) (int, error) {
// 	if num2 == 0 {
// 		return 0, utils.DibagiNolError{}
// 	} else if num1 < 0 || num2 < 0 {
// 		return 0, utils.AngkaMinus{}
// 	}
// 	result := (num1 / num2)
// 	return result, nil

// }

// func printGreeting(d dictionary.Dict) {
// 	fmt.Println(d.GetMorningGreeting())
// }
