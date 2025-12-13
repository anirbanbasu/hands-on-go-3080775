// types/structs/fields/begin/main.go
package main

import "fmt"

// define a struct type for author
type author struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	// intialize author
	//
	a := author{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
	}

	// print the author
	fmt.Printf("%#v\n", a)
}
