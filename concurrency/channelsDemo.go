package main

import "fmt"

/*
- Channels are a way for goroutines to communicate
- A channel passes a value from one goroutine to another
- "Share memory by communicating"
- Goroutines with channels are unique feature of Go language
*/

func main() {
	// create a channel of type boolean
	ch := make(chan bool)
	go waitTillChannelHasData(ch, " world ")
	fmt.Print("Hello")
	// Pass data via the channel. This will signal the waitTillChannelHasData
	// function to continue as it was blocked on the channel
	ch <- true

	// Wait until another signal is received on ch before exiting main function
	<-ch
}

func waitTillChannelHasData(channel chan bool, msg string) {
	// This call will block until the data is available in the channel
	// Once the data is available, it prints the msg that is passed-in
	if ch := <-channel; ch {
		fmt.Println(msg)
	}

	// Send another signal on channel so that the main function continues
	channel <- true
}
