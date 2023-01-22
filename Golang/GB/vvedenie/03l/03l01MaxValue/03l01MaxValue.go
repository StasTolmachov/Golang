// нахождение второго максимального числа
package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{23, 99, 5, 77, 79, 38, 45, 3, 66}
	vMax := math.MinInt64
	vMaxSecond := math.MinInt64

	for i := 0; i < len(arr); i++ {
		if vMax < arr[i] {
			vMaxSecond = vMax
			vMax = arr[i]
		} else if vMaxSecond < arr[i] {
			vMaxSecond = arr[i]
		}
	}

	fmt.Println(arr, vMax, vMaxSecond)
}
