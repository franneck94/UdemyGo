package main

import "fmt"

// Type Conversion

func main() {
	speed := 100
	force := 2.5

	// No implicit conversions! :)
	// speed = speed * force
	// speed = speed * int(force)
	speed = int(float64(speed) * force)

	fmt.Println(speed)
}
