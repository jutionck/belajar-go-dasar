package main

/*
type student struct {
	firtsName, lastName string
	age                 int
	studentAddress      address
}

type address struct {
	homeAddress map[string]string
	postCode    string
}

func main() {
	// newStudent := student{
	// 	firtsName: "budi",
	// 	lastName:  "anduk",
	// 	age:       18,
	// }

	// student1 := &newStudent
	// student2 := student1
	// student2.age = 40
	// student2.lastName = "tono"
	// student1.firtsName = "tini"

	// fmt.Println(newStudent)

	// casecase struct
	siswaBudi := student{
		firtsName: "budi",
		lastName:  "anduk",
		age:       18,
		// casecade
		studentAddress: address{
			homeAddress: map[string]string{
				"alamat1": "subang",
				"alamat2": "bandung",
			},
			postCode: "41254",
		},
	}
	fmt.Println(siswaBudi.studentAddress.homeAddress)
}
*/
// contoh dari bang edo
/*
type addressMap map[string]string

type student struct {
	firstName, lastName string
	age                 int
	address             //di ambil dari struct address, kalo ganti struct nya ganti juga ya
}
type address struct {
	homeAddress   addressMap
	officeAddress addressMap
}

func main() {
	student1 := student{
		firstName: "budi",
		lastName:  "anduk",
		age:       18,
		address: address{homeAddress: addressMap{
			"alamat1": "subang",
			"alamat2": "bandung",
		}, officeAddress: addressMap{
			"alamat1": "Bekasi",
			"alamat2": "Surabaya",
			"alamat3": "Semarang",
		},
		},
	}
	fmt.Println(student1)
}
*/

// contoh dari kang prana

import (
	"fmt"
)

type kunci string

type sliceKu []kunci

type addressMap map[kunci]sliceKu

type student struct {
	firstName, lastName kunci
	age                 int
	address
}
type address struct {
	homeAddress   addressMap
	officeAddress addressMap
}

func main() {
	student1 := student{
		firstName: "budi",
		lastName:  "anduk",
		age:       18,
		address: address{
			homeAddress: addressMap{"alamat1": []kunci{"subang"}, "alamat2": []kunci{"bandung"}}},
	}
	fmt.Println(student1)
}
