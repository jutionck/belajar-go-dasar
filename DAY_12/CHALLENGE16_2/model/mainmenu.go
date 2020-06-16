package model

type food struct {
	foodName    string
	timeServing int
}
type consument struct {
	name     string
	listFood []food
}
type bill struct {
	billName    string
	cookingTime int
}

var orderList []bill
var timeServing int
