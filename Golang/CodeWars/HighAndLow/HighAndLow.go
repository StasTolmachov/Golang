package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var str = "1 2 33 4 -5" // набор текста
	fmt.Println(str)

	fmt.Printf("%f\n", HighAndLow(str))

}

// func HighAndLow(in string) string {//me version

// 	a := strings.Split(in, " ") //string without separate
// 	b := make([]int, len(a))    // string to int
// 	for i := range a {
// 		b[i], _ = strconv.Atoi(a[i])
// 	}
// 	maxVal := math.MinInt64 // find min max
// 	minVal := math.MaxInt64
// 	for _, v := range b {
// 		if minVal > v {
// 			minVal = v
// 		}
// 		if maxVal < v {
// 			maxVal = v
// 		}
// 	}

// 	minValString := strconv.FormatInt(int64(minVal), 10)
// 	maxValString := strconv.FormatInt(int64(maxVal), 10)
// 	valueStringOut := fmt.Sprint(maxValString, " ", minValString)
// 	return valueStringOut
// }

func HighAndLow(in string) string {
	var tmpH, tmpL int
	for i, s := range strings.Fields(in) {
		n, _ := strconv.Atoi(string(s))
		if i == 0 {
			tmpH = n
			tmpL = n
		}

		if n > tmpH {
			tmpH = n
		}

		if n < tmpL {
			tmpL = n
		}
	}
	return fmt.Sprintf("%d %d", tmpH, tmpL)
}
