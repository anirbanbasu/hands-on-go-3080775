// concurrency/channel-non-blocking/begin/main.go
package main

import "time"

func main() {
	// declare a signal channel to read from
	timeChan1 := time.After(100 * time.Millisecond)
	timeChan2 := time.After(200 * time.Millisecond)
	timeChan3 := time.After(300 * time.Millisecond)
	timeChan4 := time.After(400 * time.Millisecond)

	// use a default case in select to prevent blocking
	for {
		select {
		case <-timeChan1:
			println("Received signal from timeChan1")
			return
		case <-timeChan2:
			println("Received signal from timeChan2")
			return
		case <-timeChan3:
			println("Received signal from timeChan3")
			return
		case <-timeChan4:
			println("Received signal from timeChan4")
			return
		default:
			println("No signals received yet")
			time.Sleep(250 * time.Millisecond)
		}
	}
}
