package main

import (
	"fmt"
	"unicode/utf8"
)

// "" strings - single line
// `` raw strings - multi line
// both have the same type

func main() {
	s := "carl"

	fmt.Println(s)
	fmt.Println(len(s)) // num of bytes

	s2 := "öhm"

	fmt.Println(s2)
	fmt.Println(len(s2)) // num of bytes

	fmt.Println(utf8.RuneCountInString(s2)) // runes are just like utf8 codepoints
}
