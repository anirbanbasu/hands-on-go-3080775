// concurrency/sync/begin/main.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	// given a list of names
	names := []string{"alice", "bob", "charlie", "dave", "eve", "frank", "grace", "heidi", "ivan", "judy", "karen", "leo", "mike", "nancy", "oscar", "peggy", "quinn", "ruth", "sam", "trudy", "uma", "victor", "wendy", "xander", "yvonne", "zach"}

	// initialize a map to store the length of each name
	namesMap := make(map[string]int)

	// initialize a wait group for the goroutines that will be launched
	var waitGroup sync.WaitGroup
	var mutex sync.Mutex

	// launch a goroutine for each name we want to process
	for _, name := range names {
		waitGroup.Add(1)
		go func(n string) {
			defer waitGroup.Done()
			mutex.Lock()
			defer mutex.Unlock()
			// compute the length of the name and store it in the map
			namesMap[n] = len(n)
		}(name)
	}

	// wait for all goroutines to finish
	waitGroup.Wait()

	// print the map
	fmt.Println(namesMap)
}
