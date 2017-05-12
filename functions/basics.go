package main

import "fmt"

/*
	Functions are central to Go.
	Go also supports first class functions, higher-order functions, user-defined function types,
	function literals, closures etc. which will be discussed separately
	This rich feature set supports a functional programming style in a strongly typed language.
*/
func add(a int, b int) int {
	sum := a + b
	// Requires explicit returns (i.e) it won't automatically return the value of the last expression
	return sum
}

// Observe that the variable 'sum' is specified as a part of function declaration
// And it is assigned to a+b in the definition. Also, observe that the return statement
// is not returning the sum.
func addDifferent(a, b int) (sum int) {
	sum = a + b
	return // returns sum
}

// Can return multiple values
// See the usage in the main method
func addMultiply(a, b int) (sum, product int) {
	sum = a + b
	product = a * b
	return // returns both sum and product
}

/*
Everything in Go is passed by value. That is, a function always gets a copy of the thing being passed.
For instance, passing an int value to a function makes a copy of the int and
passing a pointer value makes a copy of the pointer, but not the data it points to.
 */
func passByValue(val int) {
	val = 100
}

func passByReference(val *int) {
	*val = 100
}

// Recursive function
func fib(n int) int {
	if n <= 1 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	a := 100
	b := 10

	fmt.Println("Sum of 100, 10:", add(a, b))
	fmt.Println("Sum of 100, 10 differently :", addDifferent(a, b))

	sum, product := addMultiply(a, b)
	fmt.Println("Sum, Product of 100, 10 :", sum, product)

	val := 10
	passByValue(val)
	fmt.Println("New value of 10 after pass by value call :", val) // val should still be 10

	passByReference(&val)
	fmt.Println("New value of 10 after pass by reference call :", val) // val should be 100

	fmt.Println("Fib(10) :", fib(10))
}

