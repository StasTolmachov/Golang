// (-1, 2) --> 2 (-1 + 0 + 1 + 2 = 2)
package main

import "fmt"

func main() {
	a := -1
	b := 2
	fmt.Println(GetSum(a, b))
}

func GetSum(a, b int) int {
	var sum int
	if a < b {

		for a <= b {
			sum += a
			a++
		}
	}
	if b < a {
		for b <= a {
			sum += b
			b++
		}
	}
	return sum
}
