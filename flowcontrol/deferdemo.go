package main

import "fmt"

func main() {
	deferSimpleDemo()
	deferLIFODemo()
	nilValuePanicsDefer()

	fmt.Println(safeDivision(2, 0))
	// This will also print because safeDivision is taking care of recovering from panic
	fmt.Println(safeDivision(2, 1))

}

func deferSimpleDemo() {
	defer fmt.Println("This deferrred statement will be printed before exiting the function. \n",
		"This will be useful to clean up the resources, similar to what we do in \"finally block of Java\".")

	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()
}

func deferLIFODemo() {
	fmt.Println("\nIf you have multiple \"defer\" statements, then they will be printed Last In First Out format, as below. \n",
		"Instead of 1 to 5, 5 to 1 will be printed in for loop because of defer")

	for i := 1; i <= 5; i++ {
		defer fmt.Print(i, " ")
	}
}

func nilValuePanicsDefer() {
	var name string
	defer fmt.Println(name)

	fmt.Println("\nThis will not be printed as defer panics upfront")
}

func safeDivision(a, b int) int {
	// Comment the following anonymous function and see
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\n\nError :", r)
		}
	}()

	return a / b
}
