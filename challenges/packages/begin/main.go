// challenges/packages/begin/proverbs.go
package main

import (
	"fmt"

	"github.com/jboursiquot/go-proverbs"
)

// import the proverbs package

// getRandomProverb returns a random proverb from the proverbs package
func getRandomProverb() (saying, link string) {
	// call the appropriate function from the proverbs package to get a random proverb
	randomProverb := proverbs.Random()
	saying = randomProverb.Saying
	link = randomProverb.Link
	return
}

func main() {
	// print the result of calling your getRandomProverb function
	fmt.Println(getRandomProverb())
}
