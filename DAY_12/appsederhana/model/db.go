package model

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type dataMhs struct {
	namaMhs    string
	umurMhs    int
	jurusanMhs string
}

var mhs []dataMhs

var nama, jurusan string
var umur int
var path = "dbmhs.txt"

func totalMhs() int {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic("Failed Opened !")
	}
	str := string(b)
	// fmt.Print(str)
	return len(str)

}

func bacaFilePerBaris(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
func bacaFile(path string) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	fmt.Println(string(body))
}

func nambahTulisan(fileName string, tulisan string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer f.Close()

	if _, err = f.WriteString(tulisan); err != nil {
		log.Fatalf(err.Error())
	}

}
