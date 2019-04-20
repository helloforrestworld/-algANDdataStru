package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// rune
	s := "Yes我爱慕课网!"
	fmt.Println(len(s))
	for _, b := range []byte(s) {
		// 每个中文三字节 UTF-8
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s {
		fmt.Printf("(%d, %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("i=%d,ch=%c ", i, ch)
	}

	// 其他字符串操作strings
	//strings.Fields()
	//strings.Split()
	//strings.Join()
	// ... Contains, Index, ToLower, ToUpper
	// ... Trim, TrimRight, TrimLeft
}
