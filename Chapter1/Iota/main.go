package main

import "fmt"

func main() {
	// built-in constant generator
	const (
		mon = iota + 1
		tue
		wed
		thu
		fri
		sat
		sun
	)
	// iota will start at 0 again afterwards

	fmt.Println(mon, tue, wed, thu, fri, sat, sun)

	const (
		EST = -(5 + iota)
		_
		MST
		PST
	)

	fmt.Println(EST, MST, PST)
}
