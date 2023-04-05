package main

import "fmt"

func main() {
	mySlice := []int{11, 22, 33, 44, 55}

	for index, value := range mySlice {
		fmt.Println("index", index, "valie:", value)
	}

	fmt.Println(SquareSum2(mySlice))
}

func SquareSum2(numbers []int) (sum int) {
	for _, value := range numbers {
		sum += value * value
	}
	return sum
}
