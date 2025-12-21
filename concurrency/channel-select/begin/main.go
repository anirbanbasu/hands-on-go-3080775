// concurrency/channel-select/begin/main.go
package main

import "time"

func main() {
	// declare an empty struct channel for signaling
	signalChan := make(chan struct{})

	// declare a timer channel
	timerChan := time.After(2 * time.Second)

	// launch a goroutine to signal after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		signalChan <- struct{}{}
	}()

	// wait for a signal on either channel
	//
	select {
	case <-signalChan:
		println("Received signal from signalChan")
	case <-timerChan:
		println("Received signal from timerChan")
	}
}
