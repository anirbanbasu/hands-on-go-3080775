// generics/type-parameters/begin/main.go
package main

import "fmt"

func sumInts(a, b int) int {
	return a + b
}

func sumFloats(a, b float64) float64 {
	return a + b
}

// create generic sum function with type parameter T constrained to int and float64 types
//

func sum[T ~int | ~float64](a, b T) T {
	return a + b
}

type specialInt int
type specialFloat float64

func main() {
	// non-generic sum int function
	fmt.Println(sumInts(1, 2))

	// non-generic sum float function
	fmt.Println(sumFloats(1.3, 2.2))

	// call on generic sum function
	fmt.Println("Generic sum with ints:", sum(1, 2))
	fmt.Println("Generic sum with floats:", sum(1.3, 2.2))

	// define a compatible custom type call on generic sum function with it
	//
	one := specialInt(5)
	two := specialInt(10)
	fmt.Println("Generic sum with specialInt:", sum(one, two))

	pi := specialFloat(3.14)
	tau := specialFloat(6.28)
	fmt.Println("Generic sum with specialFloat:", sum(pi, tau))
}

// list is a singly-linked list that holds values of any type
type list[T any] struct {
	head *node[T]
	tail *node[T]
}

type node[T any] struct {
	value T
	next  *node[T]
}
