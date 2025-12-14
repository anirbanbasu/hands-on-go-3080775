// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
		}
	}()

	// use os package to read the file specified as a command line argument
	fileToRead := os.Args[1]
	bytesSlice, err := os.ReadFile(fileToRead)
	if err != nil {
		panic(fmt.Errorf("Failed to read file: %s", err))
	}

	// convert the bytes to a string
	stringData := string(bytesSlice)

	// use fmt to print out the file content
	fmt.Println("Parsed content from file", fileToRead, "as string of length", len(stringData), "characters")

	// initialize a map to store the counts
	//
	counts := make(map[string]int)

	// use the standard library utility package that can help us split the string into words
	words := strings.Fields(stringData)

	// capture the length of the words slice
	counts["words"] = len(words)

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	//
	for _, word := range words {
		for _, char := range word {
			if unicode.IsLetter(char) {
				counts["letters"]++
				if unicode.IsLower(char) {
					counts["lowercase letters"]++
				} else if unicode.IsUpper(char) {
					counts["uppercase letters"]++
				}
			} else if unicode.IsDigit(char) {
				counts["numbers"]++
			} else if unicode.IsSymbol(char) {
				counts["symbols"]++
				counts[string(char)]++
			} else if unicode.IsPunct(char) {
				counts["punctuation"]++
				counts[string(char)]++
			} else {
				counts["others"]++
			}
		}
	}

	// dump the map to the console using the spew package
	//
	spew.Dump(counts)
}
