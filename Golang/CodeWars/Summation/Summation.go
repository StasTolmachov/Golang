/*
	Write a program that finds the summation of every number from 1 to num.

The number will always be a positive integer greater than 0.
8 -> 36 (1 + 2 + 3 + 4 + 5 + 6 + 7 + 8)
*/
package main

import "fmt"

func main() {
	a := 9
	fmt.Println(Summation(a))
}

func Summation(n int) int { //my version
	sum := n
	for n >= 1 {
		sum += n - 1
		n = n - 1
	}
	return sum
}

// func Summation(n int) int {
//     return n * (n + 1) / 2
// }

// func Summation(n int) (sum int) {
// 	for i := 0; i <= n; i++ {
// 		sum += i
// 	}
// 	return
// }