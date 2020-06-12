package service

import "strconv"

// SlotMachine is struct
type SlotMachine struct {
	RandomNumber  []int
	CreditBalance int
}

// Message is Recieve method
func (s SlotMachine) Message() string {
	x := s.RandomNumber[0]
	y := s.RandomNumber[1]
	z := s.RandomNumber[2]
	getDollar := 0
	saldoCreditBalane := 0
	if x == y && y == z {
		getDollar = (x + y + z) * 200
		saldoCreditBalane = getDollar + s.CreditBalance
		return "You Win " + strconv.Itoa(getDollar) + "\nYour Total Credit Balance Is " + strconv.Itoa(saldoCreditBalane) + " Dollar"
	} else if x == y || y == z || x == z {
		getDollar = (x + y + z) * 100
		saldoCreditBalane = getDollar + s.CreditBalance
		return "You Win " + strconv.Itoa(getDollar) + "\nYour Total Credit Balance Is " + strconv.Itoa(saldoCreditBalane) + " Dollar"
	} else {
		getDollar = (x + y + z) * 50
		saldoCreditBalane = s.CreditBalance - getDollar
		return "You Lose " + strconv.Itoa(getDollar) + "\nYour Total Credit Balance Is " + strconv.Itoa(saldoCreditBalane) + " Dollar"
	}
}
