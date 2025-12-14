// flow-control/loop-basic/begin/main.go
package main

import "fmt"

func main() {
	// declare a string to iterate over
	s := "Hello world"

	// iterate over the string with basic for loop
	//
	for i := 0; i < len(s); i++ {
		// print each character
		fmt.Println(i, ":", string(s[i]))
	}
}
