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
	fmt.Println()
	withGoRoutineAnonymousFunction()
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

// Goroutines are executed within the same address space of the caller
// Interesting... This will still print " Hello world! "
func withGoRoutineAnonymousFunction() {
	word := "Hello"
	go func() {
		time.Sleep(time.Millisecond * 100)
		// word value will no longer be "Hello", rather it will be " world! "
		// because, its value was changed outside of this anonymous function below (i.e) after fmt.Print(" Hello ") statement
		fmt.Print(word)
	}()
	fmt.Print(" Hello ")
	word = " world! "
	time.Sleep(time.Millisecond * 300)
}

func printMsg(str string) {
	time.Sleep(time.Millisecond * 100)
	fmt.Print(str)
}
