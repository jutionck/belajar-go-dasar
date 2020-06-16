package model

import (
	f "fmt"
)

// PrintResult func
func PrintResult() {
	printChannel := make(chan string)
	budi := consument{"Budi", []food{{"Nasi Goreng", 4}, {"Ayam Goreng", 5}, {"Jus Alpukat", 2}}}
	tono := consument{"Tono", []food{{"Nasi Goreng", 4}}}
	jution := consument{"Jution", []food{{"Air Mineral", 1}}}
	var consumentList = []consument{budi, tono, jution}
	for _, v := range consumentList {
		for _, buy := range v.listFood {
			if buy.foodName == buy.foodName {
				temp := buy.timeServing * buy.timeServing
				timeServing += temp
				temp = 0
			}

		}
		consument := bill{
			billName:    v.name,
			cookingTime: timeServing,
		}
		orderList = append(orderList, consument)
		timeServing = 0
	}
	for _, v := range orderList {
		go cooking(v.billName, v.cookingTime, printChannel)
	}
	for i := 0; i < len(orderList); i++ {
		f.Println(<-printChannel)
	}
}
