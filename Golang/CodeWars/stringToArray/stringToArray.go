package main

import (
	"fmt"
	"strings"
)

func main() {

	text := "I love arrays they are my favorite"

	var str = strings.Split("I love arrays they are my favorite", " ")

	fmt.Printf("%q\n", str)

	fmt.Printf("%q\n", StringToArraySplit("I love arrays they are my favorite"))

	fmt.Printf("%q\n", StringToArrayFields(text))

}

// разделяет по пробелу
func StringToArrayFields(str string) (sliceString []string) {
	return strings.Fields(str)

}

// разделяет по указанному символу
func StringToArraySplit(str string) (sliceString []string) {
	return strings.Split(str, " ")

}
