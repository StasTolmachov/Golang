// игра в рулетку
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var cash, wallet, betRate, rouletteNumber, winNumber int
	betNumbersArr := [34]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33}
	var betNumberArr int

	cash = 1000
	wallet = cash

	for i := 1; wallet > 0; i++ { //начало цикла
		fmt.Println("начался цикл пока кошелек не пустой", i)
		betNumber := 55
		fmt.Printf("Wallet: %v\n", wallet)
		fmt.Printf("Input betNumber: ")
		fmt.Scanln(&betNumber)

		if betNumber == 66 { //exit
			panic(betNumber)
		}

		for betNumberArr = range betNumbersArr {
			fmt.Println("начался цикл перебора чисел ставки", betNumberArr)
			if betNumber == betNumberArr { //здесь вся игра
				fmt.Println("если число совпало с массивом")
				betRate = 0
				fmt.Printf("Input betRate: ")
				_, err := fmt.Scanf("%d", &betRate)
				if err == nil {

					if betRate < wallet {
						wallet = wallet - betRate

						rand.Seed(time.Now().UnixNano()) //рандом нахождение числа рулетки
						min := 0
						max := 33
						rouletteNumber = rand.Intn(max-min+1) + min
						fmt.Println("rouletteNumber", rouletteNumber)

						if rouletteNumber == betNumber { //проверка выиграша
							winNumber = betRate * 5
							wallet = wallet + winNumber
							fmt.Printf("You Win: %v\n", winNumber)
						}
						fmt.Println("you lost")
						break
					}
				}

			}

		}

	}

}
