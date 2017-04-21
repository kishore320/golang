package main

import "fmt"

func main() {
	// A closure is a stack frame which is allocated when a function starts its execution,
	// and not freed after the function returns
	// (as if a 'stack frame' were allocated on the heap rather than the stack!).
	increment := incrementer()
	fmt.Println(increment())
	// Should continue incrementing
	fmt.Println(increment())
	fmt.Println(increment())

	newIncrement := incrementer()
	// Should again start from 1
	fmt.Println(newIncrement())
	fmt.Println(newIncrement())

	hello := helloClosure("Bangalore")
	hello("Naga")
	hello("Kishore")
}

// A closure is created when a function returns another function that accesses the
// variables in the scope of its first function. Here 'i' is such variable
func incrementer() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// Another closure
func helloClosure(city string) func(name string) {
	return func(name string) {
		fmt.Println("Hello", name+",", "welcome to", city)
	}
}
