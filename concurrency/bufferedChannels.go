package main

import (
	"fmt"
)

/*
- By default, channels are not buffered (i.e) data can be sent via channel only if there is a receiver.
Multiple data values can not be sent via channel if there is no corresponding receiver.
- On the other hand, buffered channels can accept a limited set of values with out corresponding receivers for those values.
- Senders will block only when the buffer is full.
- Receivers will block only when the buffer is empty.
*/
func main() {
	// create a buffered channel with size 2
	ch := make(chan string, 2)
	// main function will block once two values are send via the channel
	ch <- "Hello"
	ch <- " world "

	// Initially receivers will block as buffer will be empty
	// Once the data is sent via the channel, receivers will unblock and
	// print the data in the channels
	fmt.Print(<-ch)
	fmt.Print(<-ch)
}
