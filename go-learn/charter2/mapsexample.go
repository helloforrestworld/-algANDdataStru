package main

import "fmt"

func lenOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		lastCharIndex, ok := lastOccurred[ch]
		if ok && lastCharIndex >= start {
			start = lastCharIndex + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	// 寻找最长不包含有重复字符的子串长度
	fmt.Println(
		lenOfNonRepeatingSubStr("abcabcbb"),
		lenOfNonRepeatingSubStr("bbbbb"),
		lenOfNonRepeatingSubStr("pwwkew"),
		lenOfNonRepeatingSubStr(""),
		lenOfNonRepeatingSubStr("b"),
		lenOfNonRepeatingSubStr("abcdef"),
		lenOfNonRepeatingSubStr("这里支持中文字符，字符"))
}
