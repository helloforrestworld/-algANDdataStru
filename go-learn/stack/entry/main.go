package main

import (
	"awesomeProject/stack"
	"fmt"
)

func main() {
	q := stack.Create{}
	q.Push(1)
	fmt.Println(q)
	q.Push(2)
	fmt.Println(q)
	q.Push(3)
	fmt.Println(q)
	q.Pop()
	fmt.Println(q)
	fmt.Println(q.IsEmpty())
	q.Pop()
	fmt.Println(q)
	q.Pop()
	fmt.Println(q)
	q.Pop()
	fmt.Println(q.IsEmpty())
}
