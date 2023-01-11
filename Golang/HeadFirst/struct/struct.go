package main

import "fmt"

func main() {
	type subscriber struct {
		name   string
		rate   float64
		active bool
	}
	var subscriber1 subscriber
	subscriber1.name = "Adam"
	subscriber1.rate = 4.99
	subscriber1.active = true

	var subscriber2 subscriber
	subscriber2.name = "Beth"

	fmt.Printf("%v%v\n", subscriber1, subscriber2)

}
