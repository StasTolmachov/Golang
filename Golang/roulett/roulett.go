package main

import "fmt"

func main() {
	var cash, wallet, betNumber, betRate, rouletteNumber, winNumber int

	cash = 1000
	wallet = cash

	fmt.Println(wallet)
	fmt.Scanln(&betNumber)
	fmt.Scanln(&betRate)
	wallet = wallet - betRate
	rouletteNumber = 23
	if rouletteNumber == betNumber {
		winNumber = betRate * 5
		wallet = wallet + winNumber
	}


	fmt.Println(cash, wallet, betNumber)
}
