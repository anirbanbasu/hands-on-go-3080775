// types/pointers/begin/main.go
package main

import "fmt"

func main() {
	// create a variable of type *T where T is an int
	var a *int

	// declare and assign `b` variable of type int
	var b int = 42

	// assign the address of b to a
	a = &b

	// print out the value of a which is the address of b
	fmt.Printf("address: %p\n", a)

	// print out the value at the address of b
	fmt.Printf("value: %d\n", *a)

	fmt.Printf("%d [%p]\n", *a, a)
}
