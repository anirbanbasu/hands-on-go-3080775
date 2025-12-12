// types/conversion/begin/main.go
package main

import "fmt"

func main() {
	// declare float and convert it to an unsigned int
	a := -42.575
	b := uint64(a)

	fmt.Printf("%v=%T %v=%T\n", a, a, b, b)

	a = 42.575
	b = uint64(a)

	fmt.Printf("%v=%T %v=%T\n", a, a, b, b)
}
