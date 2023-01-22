// Нахождение индексов максимального и минимального элемента массива
package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{23, 3, 5, 77, 79, 38, 45, 5, 66}
	var valueMin, valueMax, indexMin, indexMax int
	valueMin = math.MaxInt64
	valueMax = math.MinInt64
	var sum int
	fmt.Println(valueMin, valueMax)
	for i := 0; i < len(arr); i++ {
		value := arr[i]
		if value < valueMin {
			valueMin = value
			indexMin = i
		}
		if value > valueMax {
			valueMax = value
			indexMax = i
		}
	}
	fmt.Println("indexMin:", indexMin, "indexMax:", indexMax)

	for i := indexMin; i <= indexMax; i++ {
		sum += arr[i]
	}

	fmt.Println(sum)

}
