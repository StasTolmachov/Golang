// игра в рулетку
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var cash int        //начальная сумма
	var wallet int      // кошелек игрока
	var betColor string // цвет ставки
	var betNum int      // число ставки
	var betRate int     // размер ставки
	var roulettNum int  //число выпало на рулетке
	var winSum int      //сумма выиграша
	betNumbersArr := []int{0, 1, 2, 3}
	var betNumArr int //число допустимой ставки
	cash = 1000
	wallet = cash

	for i := 1; wallet > 0; i++ { //начало цикла игры со счетчиком
		fmt.Println("начался цикл пока кошелек не пустой", i)

		// цикл ввода цвета ставки
		fmt.Println("Введите цвет ставки или выход")
		for { //цикл ввода цвета ставки только допустимых значений
			fmt.Scan(&betColor)
			if betColor == "red" {
				break
			}
			if betColor == "black" {
				break
			}
			if betColor == "exit" {
				break
			}
			fmt.Println("неверный ввод")

		}
		if betColor == "exit" {
			break
		}

		fmt.Printf("Input betNum: ")
		for { //цикл ввода числа ставки допустимых значений
			fmt.Println("введите число ставки")
			for _, betNumArr = range betNumbersArr {
				if betNum == betNumArr {
					break
				}
			}
			if betNum == betNumArr {
				break
			}
			fmt.Println("неверный ввод")
		}

		fmt.Printf("Input betRate: ") //ввод размера ставки

		rand.Seed(time.Now().UnixNano()) //рандом нахождение числа рулетки
		min := 0
		max := 33
		roulettNum = rand.Intn(max-min+1) + min
		fmt.Println("roulettNum", roulettNum)

		if roulettNum == betNum { //проверка выиграша
			winSum = betRate * 2
			wallet += winSum
			fmt.Printf("You Win: %v\n", winSum)
		}

	}
	fmt.Println(betColor, betNumbersArr)
}
