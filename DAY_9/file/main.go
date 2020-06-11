package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// fmt.Println("FOO:", os.Getenv("FOO"))
	//fmt.Println("File")
	//bacaFilePerBaris("file.txt")
	config := os.Getenv("fileConfig")
	fmt.Println("File Config:")
	tulisFile(config, "file.txt", []byte("Kalau berjalan\n"))
	nambahTulisan(config, "file.txt", "prok...prok...prok\n")
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

func tulisFile(fileName string, data []byte) {
	err := ioutil.WriteFile(fileName, data, 0777)
	if err != nil {
		fmt.Println(err)
	}
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
