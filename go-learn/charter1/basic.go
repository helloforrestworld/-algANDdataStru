package main

import (
	"fmt"
	"math"
)

// 不是全局变量而是包内部变量
// var aa = 3
// 函数外面不可以:=
// bb := true
// ()批量声明
var (
	aa = "aa"
	bb = "33"
	cc = true
)

// 定义变量
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

// 变量赋值
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

// 类型推断
func variableTypeDeduction() {
	var a, b, c, d = 3, 4, true, "def"
	fmt.Println(a, b, c, d)
}

// := 省略var
func variableShorter() {
	a, b, c, d := 3, 4, true, "def"
	fmt.Println(a, b, c, d)
}

func triangle() {
	var a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 常量
func consts() {
	const (
		filename = "abc.png"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	println(filename, c)
}

// 枚举类型
func enums() {
	// const (
	// 	cpp    = 0
	// 	java   = 1
	// 	python = 2
	// 	golang = 3
	// )
	const (
		// 常量计数器iota
		cpp = iota

		python
		golang
		javascript
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, javascript, python, golang) // 0 3 1 2
}

func main() {
	// variable.go
	fmt.Println("hello, world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, cc)
	triangle()
	// const.go
	consts()
	enums()
}

// 变量定义

// 内建变量类型
// bool, string
// (u)int, (u)int8, (u)int16, (u)int32, (u)int64,  unitptr(指针)
// byte, rune
// float32, float64, complex64, complex128

// 强制类型转换
