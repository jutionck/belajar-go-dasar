package main

import "fmt"

type student struct {
	firtsName, lastName string
	age                 int
}

func main() {
	newStudent := student{
		firtsName: "budi",
		lastName:  "anduk",
		age:       18,
	}

	student1 := &newStudent
	student2 := student1
	student2.age = 40
	student2.lastName = "tono"
	student1.firtsName = "tini"

	fmt.Println(newStudent)
}
