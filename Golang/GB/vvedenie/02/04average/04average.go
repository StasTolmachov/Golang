// Найти среднее арифметическое среди всех элементов массива.
package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 33}
	var sum int
	// for _, v := range arr {
	// 	sum += v
	// }

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println(sum)
}
