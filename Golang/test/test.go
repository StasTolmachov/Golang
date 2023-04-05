package main

import (
	"fmt"
	"strconv"
)

func main() {
	minVal := -5
	maxVal := 88
	d := make([]string, 2)
	d[0] = strconv.FormatInt(int64(minVal), 10)
	d[1] = strconv.FormatInt(int64(maxVal), 10)
	fmt.Println(d)

	var str string
	str = fmt.Sprint(d[0], " ", d[1])
	fmt.Println(str)
	
}
