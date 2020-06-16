package model

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type contact struct {
	id      int
	name    string
	groupID int
}

type group struct {
	id   int
	name string
}

var (
	contacts                           []contact
	groups                             []group
	contactID, contactGroupID, groupID int
	contactName                        string
	groupName                          string
)

var path = "dbcontact.txt"
var pathGroup = "dbgroup.txt"

// total maksimal 5
func bacaFileKontak(path string) string {
	// buka file
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return ""
	}
	defer file.Close()

	// baca file
	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if err != nil {
				break
			}
		}
		if n == 0 {
			break
		}
	}
	if err != nil {
		return ""
	}

	fmt.Println("==> file ", path, "terbaca")
	fmt.Println(string(text))
	return ""
}

func totalKontak() int {
	// Catatan :
	// Udah ke baca filenya dan isinya, tapi convert nya supaya ke baca sesuai isi belum
	// var total string
	// lines, err := readLines(pathGroup)
	// if err != nil {
	// 	log.Fatalf("readLines: %s", err)
	// }
	// for _, line := range lines {
	// 	total = line
	// }
	// fmt.Println(total)
	return len(contacts)

}

// total maksimal 3
func bacaFileGroup(path string) string {
	// buka file
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return ""
	}
	defer file.Close()

	// baca file
	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if err != nil {
				break
			}
		}
		if n == 0 {
			break
		}
	}
	if err != nil {
		return ""
	}

	fmt.Println("==> file ", pathGroup, "terbaca")
	fmt.Println(string(text))
	return ""
}

func totalGroup() int {
	// Catatan
	// Sama aja nih bang

	// data := bacaFileGroup(pathGroup)
	// if data == "" {
	// 	return 0
	// }
	return len(groups)

}

func readLines(pathGroup string) ([]string, error) {
	file, err := os.Open(pathGroup)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
