package main

import "fmt"

func main() {
	// create a pointer to an int
	ptr := new(int)
	changeIt(ptr)
	fmt.Println(*ptr)

	val := 10
	changeIt(&val)
	fmt.Println(val)
}

func changeIt(ptr *int) {
	*ptr = 100
}
