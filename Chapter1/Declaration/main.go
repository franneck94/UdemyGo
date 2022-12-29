package main

import (
	"fmt"
	"path"
)

// Example: func Split(path string) (dir, file string)

func main() {
	var dir, file string // uses the default value for the dtype
	dir, file = path.Split("css/main.css")

	fmt.Println("dir: ", dir)
	fmt.Println("file: ", file)

	_, file2 := path.Split("html/main.html") // Short declaration
	fmt.Println("file2: ", file2)

	var (
		video    string
		duration int
	)
	fmt.Println("video: ", video)
	fmt.Println("duration: ", duration)
}
