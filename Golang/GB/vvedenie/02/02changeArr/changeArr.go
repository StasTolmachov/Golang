//Задача 2. Задание на «разворот» массива. Нужно перевернуть массив и записать его в обратном порядке.
package main

import "fmt"

func main() {
	arr1 := []int{11, 22, 33, 44, 55}
	var arr2 []int
	arr2 = append(arr2, arr1...)
	var i1 int
	i1 = len(arr1) - 1
	var i2 int

	for i1 >= 0 && i2 < len(arr2) {
		arr2[i2] = arr1[i1]
		i1 = i1 - 1
		i2 = i2 + 1
	}
	arr1 = arr2
	fmt.Println(arr1)

}
