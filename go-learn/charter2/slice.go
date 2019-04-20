package main

import "fmt"

func updateArr(arr []int) {
	arr[0] = 100
}

// slice是array的视图

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:5] = ", arr[:5])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])

	ordinary := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("originArr is ", ordinary)
	s1 := ordinary[1:]
	s2 := ordinary[:]
	fmt.Println("s1 is ", s1)
	fmt.Println("s2 is ", s2)

	fmt.Println("After updateArr call")
	updateArr(s1)
	updateArr(s2)

	fmt.Println("s1 is ", s1)
	fmt.Println("s2 is ", s2)
	fmt.Println("ordinary is ", ordinary)

	fmt.Println("reslice")
	fmt.Println("ordinary is ", ordinary)
	reslice1 := ordinary[2:6]
	fmt.Println("reslice1 is", reslice1)
	reslice2 := reslice1[0:2]
	fmt.Println("reslice2 is", reslice2)

	ordinary[0] = 0
	ordinary[1] = 1
	fmt.Println("init ordinary, and it go back to: ", ordinary) // [0, 1, 2, 3, 4, 5, 6, 7]
	fmt.Println("test slice extend")
	s1 = ordinary[2:6] // [2,3,4,5]
	s2 = s1[3:5]       // 报错？
	fmt.Println("s1 is:", s1)
	fmt.Println("s2 is", s2) // [5, 6]

	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",
		s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n",
		s2, len(s2), cap(s2))

	fmt.Println("append to slice")
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3, s4, s5 = ", s3, s4, s5)
	fmt.Println("ordinary = ", ordinary)

	// 添加元素如果超过cap, 系统会重新分配更大的底层数组
}
