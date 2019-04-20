package main

import "fmt"

// Go语言不能运算
// Go只有值传递，没有引用传递

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	var a int = 2
	var pa *int = &a
	*pa = 10
	fmt.Println(a)

	numa, numb := 1, 2
	swap(&numa, &numb)
	fmt.Println(numa, numb)
}
