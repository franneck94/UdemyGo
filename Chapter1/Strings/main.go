package main

import "fmt"

// "" strings - single line
// `` raw strings - multi line
// both have the same type

func main() {
	var s string

	s = "how are you?"
	fmt.Println(s)

	s = `
how
are
you?
`
	fmt.Println(s)

	s = "<html>\n\t<body>\"Hello</body>\n</html>" // interpretes \n and \t
	fmt.Println(s)

	s = `<html>\n\t<body>\"Hello</body>\n</html>` // just text
	fmt.Println(s)

	s = `
<html>
	<body>\"Hello</body>
</html>` // just text
	fmt.Println(s)
}
