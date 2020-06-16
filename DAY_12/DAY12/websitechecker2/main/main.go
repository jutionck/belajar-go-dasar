package main

import (
	"fmt"
	"net/http"
	"time"
)

// var wg sync.WaitGroup

func init() {
	fmt.Println(" Go Routine ")
	fmt.Println(" =========== ")
}

// func main merupakan salah satu goroutine
func main() {
	links := []string{
		"http://golang.org",
		"http://facebook.com",
		"http://www.kompas.com",
		"http://google.com",
	}
	c := make(chan string)
	fmt.Printf("Start : %v\n", time.Now())

	for _, link := range links {
		// penggunaan Go untuk pengecekan (concorency)
		// cek satu-satu ya, yang selesai keluar hasilnya
		go checkLink(link, c)
		// jika tanpa go jadi sequintial (proses pengerjaan di waktu yang sama)
	}
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
	fmt.Printf("End : %v\n", time.Now())
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Mungkin bermasalah")
		c <- ""
		return
	}
	c <- "Ok"
	fmt.Println(link, "Ok!!!")
}
