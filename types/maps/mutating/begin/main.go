// types/maps/mutating/begin/main.go
package main

import "fmt"

type author struct {
	name string
}

func main() {
	var authors = map[string]author{
		"tm": {name: "Toni Morrison"},
		"ma": {name: "Marcus Aurelius"},
		"rt": {name: "Rabindranath Tagore"},
	}
	fmt.Printf("%v\n", authors)

	// change the value of an author in the map
	authors["ma"] = author{name: "Michael Madhusudan Dutta"}

	// print the map
	fmt.Printf("%v\n", authors)

	// delete a key from the map
	delete(authors, "ma")

	// print the map
	fmt.Printf("%v\n", authors)
}
