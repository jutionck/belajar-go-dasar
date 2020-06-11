package main

import "fmt"

type appError struct {
	customMessage string
}

func (a appError) Error() string {
	return a.customMessage
}

// Constructor Custom Error
func newAppError(message string) *appError {
	return &appError{customMessage: message}
}

func main() {
	totalGaji, err := salaryCalculation()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Total Gaji Bulan ini", totalGaji)
}

func salaryCalculation() (t float64, e error) {
	var totalGajiTemp float64
	var aErr error = nil
	defer func() {
		if err := recover(); err != nil {
			e = newAppError(fmt.Sprintf("%v", err))
			t = totalGajiTemp
		}
	}()
	for i := 0; i <= 10; i++ {
		if i == 3 {
			panic("Fail Database")
		}
		totalGajiTemp++
	}
	return totalGajiTemp, aErr
}
