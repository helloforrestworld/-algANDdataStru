package main

import (
	"fmt"
	"io/ioutil"
)

// if条件语句
func useif() {
	const filename = "abc.txt"
	// contents, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%s \n", contents)
	// }

	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s \n", contents)
	}
}

// switch 条件语句
func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)
	}
	return result
}

func useSwitch() {
	var r1 = eval(1, 2, "+")
	var r2 = eval(1, 2, "-")
	var r3 = eval(1, 2, "*")
	var r4 = eval(1, 2, "/")
	fmt.Println(r1, r2, r3, r4)
}

func main() {
	useif()
	useSwitch()
}
