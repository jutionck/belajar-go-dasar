// [CHALLENGE - 16.1] Change Me!
/*
Age didapatkan dari tahun sekarang di kurang tahun lahir dan lakukan sebuah pengecekan
apabila tahun lahir tidak diisi atau tahun lahir lebih besar dibandingkan tahun sekarang
maka age akan berisi 'Invalid Birth Year'.
*/
// Jution Candra Kirana

package main

import (
	"fmt"
	"strconv"
	"time"
	//ambil waktu sekarang
)

type dataPersons struct {
	firstName string
	lastName  string
	gender    string
	age       string
}

var (
	now             = time.Now()
	currentYear int = now.Year()
	dbTemp      dataPersons
)

func validateAge(inputData [][]string) (dbPerson []dataPersons) {
	for i := range inputData {
		intYear, _ := strconv.Atoi(inputData[i][3]) //convert to string (tahun)
		ageTemp := currentYear - intYear
		if intYear > currentYear {
			ageString := "Invalid Birth Year"
			dbTemp = dataPersons{
				firstName: inputData[i][0],
				lastName:  inputData[i][1],
				gender:    inputData[i][2],
				age:       ageString,
			}
			dbPerson = append(dbPerson, dbTemp) //input to struct dataPerson
			ageTemp = 0
		} else {
			ageString := strconv.Itoa(ageTemp) //convert to string
			dbTemp = dataPersons{
				firstName: inputData[i][0],
				lastName:  inputData[i][1],
				gender:    inputData[i][2],
				age:       ageString,
			}
			dbPerson = append(dbPerson, dbTemp) //input to struct dataPerson
			ageTemp = 0
		}
	}
	return
}

func main() {
	var dataPerson = [][]string{
		{"Budi", "Tono", "L", "1993"},
		{"Tini", "Nurlela", "P", "2021"},
		{"Jution", "Candra", "P", "1995"},
		{"Destry", "Faradila", "P", "1998"},
	}
	changeMe := validateAge(dataPerson)
	fmt.Println(changeMe)
}
