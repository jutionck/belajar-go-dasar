package model

import (
	f "fmt"
	t "time"
)

// cooking function (GO)
func cooking(consumnet string, timeServing int, printChannel chan string) {
	// calculate the order time
	t.Sleep(t.Second * t.Duration(timeServing))
	f.Println("Enjoy your food... ", consumnet, "Time Serving :", timeServing, "Seconds")
	printChannel <- "Please pay, after eat..."
}
