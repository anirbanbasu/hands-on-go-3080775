// interfaces/type-switch/begin/main.go
package main

import "fmt"

// define whatAmI which takes in an argument of any type and returns inforamtion about the underlying value's type

func whatAmI(i any) string {
	// use a type switch to determine the underlying type of i
	switch v := i.(type) {
	case int:
		return fmt.Sprintf("I am an int with value %d", v)
	case string:
		return fmt.Sprintf("I am a string with value %q", v)
	case bool:
		return fmt.Sprintf("I am a bool with value %t", v)
	default:
		return fmt.Sprintf("I (%v) am of a different type: %T", v, v)
	}
}

func main() {
	// invoke whatAmI function
	fmt.Println(whatAmI(1))
	fmt.Println(whatAmI("hello"))
	fmt.Println(whatAmI(true))
	fmt.Println(whatAmI(3.14))
}
