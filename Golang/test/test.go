package main

import "fmt"

func main() {
	var betNum int
	wallet := 20
	for wallet > 0 { //цикл всей игры
		fmt.Println("начало цикла")
		for { //цикл проверки ввода данных
			_, err := fmt.Scanf("%d", &betNum)
			wallet -= betNum
			if err == nil {
				break
			}
		}
		fmt.Println("play")
	}
	fmt.Println("выход с цыкла игры", betNum, wallet)
}

// fmt.Println("неправельный ввод")
