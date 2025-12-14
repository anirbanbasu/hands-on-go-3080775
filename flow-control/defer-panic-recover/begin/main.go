// flow-control/defer-panic-recover/begin/main.go
package main

import "fmt"

func cleanup(msg string) {
	fmt.Println(msg)
}

func main() {
	// defer function calls
	defer cleanup("First defer")
	defer cleanup("Second defer")
	defer cleanup("Third defer")

	fmt.Println("Working in main function...")

	// defer recovery
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
		}
	}()

	// panic
	panic("Something went wrong!")
}
