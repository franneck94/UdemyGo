package main

import (
	"fmt"
	"os"
)

// A Slice can store multiple values

func main() {
	// var Args []string // slice of strings

	fmt.Printf("%#v\n", os.Args)

	fmt.Println("Argc: ", len(os.Args))
	fmt.Println("1st: ", os.Args[1])
	fmt.Println("2nd: ", os.Args[2])
}
