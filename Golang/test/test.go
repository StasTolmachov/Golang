package main

import (
	"fmt"
)

func main() {
	// var betNum int
	// wallet := 20
	// for wallet > 0 { //цикл всей игры
	// 	fmt.Println("начало цикла")
	// 	for { //цикл проверки ввода данных
	// 		_, err := fmt.Scanf("%d", &betNum)
	// 		wallet -= betNum
	// 		if err == nil {
	// 			break
	// 		}
	// 	}
	// 	fmt.Println("play")
	// }
	// fmt.Println("выход с цыкла игры", betNum, wallet)

	betNumbersArr := []int{0, 1, 2}
	var betNum int // число ставки
	var betNumArr int
	var temp int
	// betNum = 1

	for {
		fmt.Println("введите число ставки")
		fmt.Scanln(&betNum)
		// fmt.Scanf("%d", &betNum)
		// if err != nil {
		// 	fmt.Println("неправельный ввод err", err)

		// 	continue
		// }
		for _, betNumArr = range betNumbersArr {
			if betNum == betNumArr {
				break
			}
		}
		if betNum == betNumArr {
			break
		}
		fmt.Println("неправельный ввод")
	}

	fmt.Println(betNum, betNumArr, temp)
}

// fmt.Println("неправельный ввод")
// betNumbersArr := [34]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33}
