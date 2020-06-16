package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

// var wg sync.WaitGroup

func init() {
	fmt.Println(" Go Routine ")
	fmt.Println(" =========== ")
}

// func main merupakan salah satu goroutine
func main() {
	var wg sync.WaitGroup
	links := []string{
		"http://golang.org",
		"http://facebook.com",
		"http://www.kompas.com",
		"http://google.com",
	}
	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("Routine\t", runtime.NumGoroutine())
	fmt.Printf("Start : %v\n", time.Now())

	for _, link := range links {
		// penggunaan Go untuk pengecekan (concorency)
		// cek satu-satu ya, yang selesai keluar hasilnya
		wg.Add(1)
		go checkLink(link, &wg)
		// jika tanpa go jadi sequintial (proses pengerjaan di waktu yang sama)
	}
	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("Routine\t", runtime.NumGoroutine())
	fmt.Printf("End : %v\n", time.Now())
	wg.Wait()
}

func checkLink(link string, wg *sync.WaitGroup) {
	_, err := http.Get(link)
	// anonimus function (IIFE)
	// defer func() {
	// 	wg.Done()
	// }()
	defer doneNotify(wg)
	if err != nil {
		fmt.Println(link, "Mungkin bermasalah")
		return
	}
	fmt.Println(link, "Ok!!!")
}

func doneNotify(wg *sync.WaitGroup) {
	wg.Done()
}
