package main

import "fmt"

func main() {
	numberSLice := []int{1, 2, 3, 4, 5} //1 + (2*2) + (3*3) + (4*4) + (5*5) = 1+4+9+16+25 = 55
	fmt.Println(SquareSum(numberSLice))
}

func SquareSum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number * number
	}
	return sum
}
