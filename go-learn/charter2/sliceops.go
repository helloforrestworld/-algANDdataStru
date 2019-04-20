package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
}

func main() {
	var s []int // zero value for slice is nil
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	// 创建切片
	s1 := []int{2, 4, 6, 8}  // [2, 4, 6, 8]
	s2 := make([]int, 8)     // len = 8 cap  = 8
	s3 := make([]int, 8, 32) // len = 8 cap = 32
	printSlice(s2)
	fmt.Println("s1, s2, s3 = ", s1, s2, s3)

	// copy slice
	copy(s2, s1)
	fmt.Println("after copy, s2 is ", s2)
	printSlice(s2)

	// delete element
	s2 = append(s2[:3], s2[4:]...)
	fmt.Println("after delte s2 is ", s2)
	printSlice(s2)

	// pop element from front
	fmt.Println("popping from front")
	front := s2[0]
	s2 = s2[1:]
	// pop element from back
	fmt.Println("popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]

	fmt.Println("front, tail is ", front, tail)
	fmt.Println("after popping s2 is ", s2)
	printSlice(s2)
}
