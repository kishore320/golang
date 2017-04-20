package main

import (
	"fmt"
	"time"
)

// Demonstrates the use of 'go' goroutine for concurrent execution
func main() {
	withoutGoRoutine()
	fmt.Println()
	withGoRoutine()
}

// Should print " world! Hello "
func withoutGoRoutine() {
	printMsg("world!")
	fmt.Print(" Hello ")
	time.Sleep(time.Millisecond * 300)
}

// Should print " Hello world! "
func withGoRoutine() {
	go printMsg("world!")
	fmt.Print(" Hello ")
	time.Sleep(time.Millisecond * 300)
}

func printMsg(str string) {
	time.Sleep(time.Millisecond * 100)
	fmt.Print(str)
}
