package main

import (
	"fmt"
)

// 数组是值类型

func definedArr() {
	var arr1 [5]int
	arr2 := [5]int{1, 2, 3, 4, 5}
	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3, grid)
}

func eachArr() {
	arr3 := [5]int{1, 2, 3, 4, 5}

	//for i := 0; i < len(arr3); i++ {
	//	fmt.Println(arr3[i])
	//}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}
}

func main() {
	definedArr()
	eachArr()
}
