// flow-control/loop-range/begin/main.go
package main

import "fmt"

func main() {
	// initialize a slice of ints
	nums := []int{10, 20, 30, 40, 50}

	// use for-range to iterate over the slice and print each value
	for idx, num := range nums {
		fmt.Println(idx, ":", num)
	}

	// declare a map of strings to ints
	//
	m := map[string]int{
		"apple":   5,
		"banana":  10,
		"cherry":  15,
		"avocado": 20,
		"date":    25,
		"fig":     30,
	}

	// use for-range to iterate over the map and print each key/value pair
	//
	for key, value := range m {
		fmt.Println(key, ":", value)
	}
}
