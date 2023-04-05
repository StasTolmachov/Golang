// "Hello, <name> how are you doing today?"
package main

import "fmt"

func main() {
	fmt.Println(Greet("stas"))
}

func Greet(name string) (string, string, string) {
	fmt.Scanln(&name)
	return "Hello,", name, "how are you doing today?"
}
