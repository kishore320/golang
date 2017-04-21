package main

import "fmt"

func main() {
	demoPanic()
	fmt.Println("Printing after panicking.....")
}

func demoPanic() {
	// Recover from panic. Comment the following and see
	defer func() {
		recover()
	}()

	panic("PANIC !!!")
}
