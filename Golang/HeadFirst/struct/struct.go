package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Subscriber struct {
	name   string
	rate   float64
	active bool
}

func main() {

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	var subscriber1 Subscriber
	subscriber1.name = "Adam"
	subscriber1.rate = 4.99
	subscriber1.active = true

	var subscriber2 Subscriber
	subscriber2.name = "Beth"

	fmt.Printf("%v%v\n", subscriber1, subscriber2)

	subscriber5 := DefaultSubscriber("Jhon Smith")
	fmt.Println(subscriber5)

	PrintInfo(subscriber5)

}

func DefaultSubscriber(name string) *Subscriber {
	var s Subscriber
	s.name = name
	s.rate = 4.99
	s.active = true
	return &s
}
func PrintInfo(s *Subscriber) {
	fmt.Println(s.name)
	fmt.Println(s.rate)
	fmt.Println(s.active)
}
