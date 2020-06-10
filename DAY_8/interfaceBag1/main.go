package main

import "fmt"

type dictLanguage interface {
	GetMorningGreeting() string
}

type englishDict struct {
}

type japanDict struct {
}

func (en englishDict) GetMorningGreeting() string {
	return "Hello, good morning"
}

func (jp japanDict) GetMorningGreeting() string {
	return "おはようございまず"
}

func printingGreeting(d dictLanguage) {
	fmt.Println(d.GetMorningGreeting())
}
func main() {
	enDict := englishDict{}
	jpDict := japanDict{}
	printingGreeting(enDict)
	printingGreeting(jpDict)
}
