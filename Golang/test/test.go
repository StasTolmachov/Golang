package main

import (
	"fmt"
)

func main() {
	betNumber := 23
	betNumbersArr := [34]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33}
	var betNumberArr int

	// for i := 0; i < len(betNumberArr); i++ {
	// 	fmt.Println("df")
	// }

	for betNumberArr = range betNumbersArr {
		if betNumber == betNumberArr {
			fmt.Println("game start")
		}
	}
	// fmt.Printf("%v\n%v\n", betNumberArr, len(betNumberArr))

	var test int

	fmt.Println(test)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		panic(err)
	}

	fmt.Println(i, err)

}
