package main

import "fmt"

func main() {
	a := 100
	b := 10

	fmt.Println("Sum of 100, 10:", add(a, b))
	fmt.Println("Sum of 100, 10 differently :", addDifferent(a, b))

	sum, product := addMultiply(a, b)
	fmt.Println("Sum, Product of 100, 10 :", sum, product)

	res := applyLambda(a, b, func(c, d int) int {
		return c - d
	})

	fmt.Println("Substraction 10 from 100 using lambda :", res)

	val := 10
	passByValue(val)
	fmt.Println("New value of 10 after pass by value call :", val) // val should still be 10

	passByReference(&val)
	fmt.Println("New value of 10 after pass by reference call :", val) // val should be 100
}

func add(a, b int) int {
	return a + b
}

// Observe that the variable 'sum' is specified as a part of function declaration
// And it is assigned to a+b in the definition. Also, observe that the return statement
// is not returning the sum.
func addDifferent(a, b int) (sum int) {
	sum = a + b
	return // returns sum
}

// Can return multiple values
func addMultiply(a, b int) (sum, product int) {
	sum = a + b
	product = a * b
	return // returns both sum and product
}

// Pass in lambda which will be applied to the arguments a,b
func applyLambda(a, b int, fn func(c, d int) int) int {
	return fn(a, b)
}

func passByValue(val int) {
	val = 100
}

func passByReference(val *int) {
	*val = 100
}
