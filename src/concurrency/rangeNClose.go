package main

import (
	"fmt"
	"time"
)

/*
- Range can be used to receive values from a channel in a loop.
- Close is used to indicate that the channel is retired. Closing a closed channel can cause runtime panic.
*/

func main() {
	ch := make(chan int)
	go fillTheChannel(ch, 5)
	// Receive the values that are coming out of the channel using 'range'
	for val := range ch {
		fmt.Print(val, " ")
	}

	// Check the channel status
	val, ok := <-ch
	fmt.Println("\nIs channel open ? ", ok)
	fmt.Println("Value in the channel : ", val)
}

func fillTheChannel(ch chan int, count int) {
	// Send the values via the channel in a for loop
	for i := 0; i < count; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 200)
	}
	close(ch)
}
