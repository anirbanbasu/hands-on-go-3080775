// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val = "global"

func main() {
	// create a local variable "val" and assign it an int value
	val := 42

	// print the value of the local variable "val"
	fmt.Printf("%T, local val = %v\n", val, val)

	// print the value of the package-level variable "val"
	printGlobalVal()

	// update the package-level variable "val" and print it again
	updateGlobalVal("updated global")
	printGlobalVal()

	// print the pointer address of the local variable "val"
	p := &val
	fmt.Printf("%T, local &val = %v\n", p, p)

	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	*p = 100
	// Instead of setting p := &val and derefencing it again, we also use *(&val) = 100
	fmt.Printf("local val = %v\n", *p)
}

func printGlobalVal() {
	// print the value of the package-level variable "val"
	fmt.Printf("global val = %v\n", val)
}

func updateGlobalVal(newVal string) {
	// update the package-level variable "val" with the value passed as an argument
	val = newVal
}
