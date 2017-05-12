package main

import (
	"fmt"
	"github.com/kishore320/golang/utils"
)

func regularFunction(msg string) {
	fmt.Println("This message is printed in regular function")
}

// Functions with out any names
func anonymousFunctions() {
	utils.Header("Anonymous Functions")

	func(msg string) {
		fmt.Println(msg)
	}("This message is printed using anonymous function")

	// Alternatively
	anonymousFn := func(msg string) {
		fmt.Println(msg)
	}

	anonymousFn("This message is also printed using anonymous function, but bit differently")
	anonymousFn("Another message using the same anonymous function")
}

// Demonstrate the scope of local variables
func inGeneralLocalVariableScope() {
	outerVariable := 100
	for i := 1; i < 2; i++ {
		fmt.Println("Outer variable is with in the scope and available here :", outerVariable)
	}

	for i := 1; i < 2; i++ {
		innerVariable := 200
		fmt.Println(innerVariable)
	}

	 //fmt.Println(innerVariable) // innerVariable is not with in the scope and not available here
}

func localVariableScopeOutsideOfFunction() {
	// invoke the function
	inGeneralLocalVariableScope()

	// outerVariable, that is defined in inGeneralLocalVariableScope is out of scope here
	//fmt.Println(outerVariable) // Does not work
}

// Closure is a persistent scope which holds on to local variables even after the code execution
// has moved out of that scope

// A closure is created when a function returns another function that accesses the
// variables in the scope of its first function. Here 'i' is such variable
func incrementer() func() int {
	i := 0
	return func() int {
		// Outer variable i's scope will be closed with in the inner function
		// so, it will be persisted even after outer function exits
		i += 1
		return i
	}
}

func incrementerDemo() {
	increment := incrementer()

	fmt.Println("Initial Value :", increment())
	// Should continue incrementing
	// If closures are not supported, variable 'i' in incrementer() might have been garbage collected and
	// a next call to increment() would have thrown error as 'i' no longer exists. This is not the case here.
	fmt.Println("Value should get incremented :", increment())
	fmt.Println("Value should keep on incrementing :", increment())

	newIncrement := incrementer()
	// Should again start from 1
	fmt.Println("Start over again :", newIncrement())
	fmt.Println("And start incrementing again :", newIncrement())
}

// Creating function factories using closures
func employeeFactory(typeOfEmployee string) func(name string) {
	return func(employeeName string) {
		fmt.Println(typeOfEmployee, employeeName)
	}
}

func employeeFactoryDemo() {
	fmt.Println()
	managers := employeeFactory("Manager")
	managers("Naga")

	developers := employeeFactory("Developer")
	developers("Kishore")
	developers("Sidde")
}

// Simulating object oriented programming using closures
// Encapsulation (Data Hiding)
func counter() (func() int, func() int, func() int) {
	i := 0

	increment := func() int {
		i += 1
		return i
	}

	decrement := func() int {
		i -= 1
		return i
	}

	reset := func() int {
		i = 0
		return i
	}

	return increment, decrement, reset
}

func oopDemo() {
	inc, dec, reset := counter()
	fmt.Println("Incrementing :", inc())
	fmt.Println("Incrementing :", inc())
	fmt.Println("Decrementing :", dec())
	fmt.Println("Resetting :", reset())
}

func closuresDemo() {
	utils.Header("Closures")

	//incrementerDemo()
	//employeeFactoryDemo()
	oopDemo()
}

func main() {
	//localVariableScopeOutsideOfFunction()
	//anonymousFunctions()
	closuresDemo()
}


