package main // package main means, it is an executable not a library

import (
	"../utils"
	"fmt"
) // import the required packages

// single line comment
/*
Multi-line
comment
*/

// main function is called during execution
func main() {
	helloWorld()
	declareAndUseVariables()
	arithmeticaOperations()
	logicalOperations()
}

func helloWorld() {
	fmt.Println("Hello GO world!!") // Use "godoc fmt Println" from commandline to view the docs
}

func declareAndUseVariables() {
	/*
		One way to declare and assign a variable
		var <var_name> <var_type> = <var_value>
	*/
	var age int = 40
	// you can use int32/int64
	// If nothing is specified it is automatically derived based on whether machine is 32bit/64bit

	var rateOfInterest float64 = 6.9 // Semicolon is optional and generally not used

	/*
		Another way to declare and assign a variable
		No need to explicitly specify the type, runtime can automatically derive it
	*/
	havingGoodDay := true

	/*
		Following is not allowed as the language is statically typed
		havingGoodDay is a boolean (bool) and can not be assigned to string
	*/
	// havingGoodDay = "false"

	// Yet another way to create variables
	var (
		isTrue bool
		myName string
	)
	isTrue = false
	myName = "Nagakishore Sidde"

	const pi float64 = 3.14

	utils.Header("Declaration and Use of Variables")
	fmt.Printf("Age = %d, ROI = %.2f, Having GoodDay? = %t, is it True? = %t, MyName = %s, Pi = %.2f \n",
		age, rateOfInterest, havingGoodDay, isTrue, myName, pi)

	// To know the datatype, try the following
	fmt.Printf("Datatype of Age is '%T' \n", age)

	fmt.Printf("Binary representation of 5 : '%b' \n", 5)
	fmt.Printf("Character code for 65 is : '%c' \n", 65)
	fmt.Printf("Octal representation of 15 : '%o' \n", 15)
	fmt.Printf("Hexadecimal representation of 19 : '%x' \n", 19)
	fmt.Printf("Scientific notation of pi value is : '%e'\n", pi)
}

func arithmeticaOperations() {
	utils.Header("Arithmetic Operations")
	fmt.Println("8 + 2 = ", 8+2)
	fmt.Println("8 - 2 = ", 8-2)
	fmt.Println("8 * 2 = ", 8*2)
	fmt.Println("8 / 2 = ", 8/2)
	fmt.Println("8 % 2 = ", 8%2)
}

func logicalOperations() {
	utils.Header("Logical Operations")
	fmt.Println("true && false = ", true && false)
	fmt.Println("true || false = ", true || false)
	fmt.Println("!true = ", !true)
}
