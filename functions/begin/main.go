// functions/begin/main.go
package main

import (
	"errors"
	"fmt"
	"math"
)

// simple greet function
//
func greet() string {
	return "Hello, Gopher!"
}

// greetWithName returns a greeting with the name
//

func greetWithName(name string) string {
	return "Hello, " + name + "!"
}

// greetWithName returns a greeting with the name and age of the person
//

func greetWithNameAndAge(name string, age int) (greeting string) {
	greeting = "Hello, " + name + "! You are " + fmt.Sprint(age) + " years old.	"
	return
}

// divide divides two numbers and returns the result
// if the second number is zero, it returns  error
//

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return math.Inf(int(a)), errors.New(fmt.Sprint("cannot divide ", a, " by zero"))
	}
	return a / b, nil
}

func main() {
	// invoke greet function
	fmt.Println(greet())

	// invoke greetWithName function
	fmt.Println(greetWithName("Toni"))

	// invoke greetWithNameAndAge function
	fmt.Println(greetWithNameAndAge("James", 30))

	// invoke divide function
	fmt.Println(divide(10, 2))

	// invoke divide function with zero denominator to get an error
	fmt.Println(divide(5, 0))
}
