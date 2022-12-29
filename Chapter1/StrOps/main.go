package main

import (
	"fmt"
	"os"
	"strings"
)

// "" strings - single line
// `` raw strings - multi line
// both have the same type

func main() {
	msg := os.Args[1]
	len := len(msg)

	s := strings.Repeat("!", len) + strings.ToUpper(msg) + strings.Repeat("!", len)
	fmt.Println(s)
}
