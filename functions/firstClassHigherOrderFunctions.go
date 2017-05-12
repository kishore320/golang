package main

import (
	"fmt"
	"github.com/kishore320/golang/utils"
)

func firstClassFunctionsDemo() {
	utils.Header("First class functions")

	fmt.Println("Go supports first-class functions (i.e) it treats functions as values")
	fmt.Println("You can assign a function to a variable, pass it around etc.")

	// Demonstrates assigning of functions to the variables
	addFunction := func(a, b int) int {
		return a+b
	}

	substractFunction := func(a, b int) int {
		return a-b
	}

	// Demonstrates storing of functions in the data structures
	functionsArray := [2]func(int, int) int {addFunction, substractFunction}

	for _,function := range functionsArray {
		fmt.Println(function(10, 20))
	}
}

// Higher order function which takes another function as argument
// Takes two integers and a function (function which needs two integers and returns an integer)
// and applies the given function to those two integers
func calculate(a, b int, functionToApply func(c, d int) int) int {
	return functionToApply(a, b)
}

// Higher order function which returns the function as its result
// It returns a function which computes a square of given value
func getSqFunction() func(a int) int {
	return func(a int) int  {
		return a * a
	}
}

func higherOrderFunctions() {
	utils.Header("Higher order functions")

	fmt.Println("Go supports higher-order functions (i.e) functions in Go can work on other functions")
	fmt.Println("Functions in Go can take one (or) more functions as arguments")

	multiplicationFunction := func(a, b int) int {
		return a * b
	}

	divisionFunction := func(a, b int) int {
		return a / b
	}

	fmt.Println("Product of 10 and 20 :", calculate(10, 20, multiplicationFunction))
	fmt.Println("Division of 10 by 20 :", calculate(10, 20, divisionFunction))

	fmt.Println("Functions in Go can return functions too")
	sqFunction := getSqFunction()
	fmt.Println("Square of 10:", sqFunction(10))

}

func main() {
	firstClassFunctionsDemo()
	//higherOrderFunctions()
	// Note that first-class and higher-order functions bit different even though
	// they are closely related
}