package main

import "fmt"


var Bus1 = [][]int{
	{10, 0},
	{3, 5},
	{5, 8},
}

func main() {

	fmt.Println(Number(Bus1))
}

func Number(stops [][]int) (InBus int) {
	for _, pass := range stops {
		InBus += pass[0] - pass[1]

	}
	return InBus
}
