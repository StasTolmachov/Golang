package main

import "fmt"

func main() {
	var betColor string
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
	fmt.Println(betColor)
}
