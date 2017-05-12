package main // package main means, it is an executable not a library

import (
	"../utils"
	"fmt"
	"math/cmplx"
	"unsafe"
) // import the required packages

// single line comment

/*
Multi-line
comment
*/

// main function is called during execution
func main() {
	//helloWorld()

	//declareAndUseVariables()

	 //keywords()

	//dataTypes()

	arithmeticOperations()

	//comparisionOperations()

	//logicalOperations()
}

func helloWorld() {
	fmt.Println("Hello GO world!!") // Use "godoc fmt Println" from commandline to view the docs
}

func declareAndUseVariables() {
	/*
		One way to declare and assign a variable
		var <var_name> <var_type> = <var_value>
	*/
	var age int = 40;

	// var age int // declaration with out initialization
	// age = 40

	// Line separator is a statement terminator
	// So no need of ';' at the end of statement and generally not used

	/*
		Another way to declare and assign a variable
		No need to explicitly specify the type, runtime can automatically derive it
	*/
	havingGoodDay := true // Preferred way if you know the value upfront

	/*
		Following is not allowed as the language is statically typed
		havingGoodDay is a boolean (bool) and can not be assigned to string
	*/
	 //havingGoodDay = "false"

	// Yet another way to create variables
	var (
		isTrue bool
		myName string
	)
	isTrue = false
	myName = "Nagakishore Sidde"

	// Go is case-sensitive
	name := "name"
	Name := "Name"
	fmt.Println("\nGo is case-sensitive. So 'name' is different from 'Name'", name, Name)

	fmt.Printf("Age = %d, Having GoodDay? = %t, is it True? = %t, MyName = %s\n",
		age, havingGoodDay, isTrue, myName)
}

func dataTypes() {
	numericTypes()
	nonNumericDataTypes()
	dataTypesFormatting()
}

func numericTypes() {
	utils.Header("Numeric Data Types")

	// Note that byte is an alias for uint8
	var byteType byte
	byteType = 5

	// you can use int8/int16/int32/int64
	// If nothing is specified it is automatically derived based on whether machine is 8bit/16bit/32bit/64bit
	var intType int
	intType = -25456


	// you can use uint8/uint16/uint32/uint64
	var unsignedInt uint
	unsignedInt = 25456

	// you can use float64 also
	var floatType float32
	floatType = 6.75

	// you can use complex64 also
	var complexType complex128
	complexType = cmplx.Sqrt(-2 + 12i)

	// Special type for pointer. No need to assume that pointer value would always fit inside an unsigned int
	// uintptr is an integer type that is large enough to hold the bit pattern of any pointer.
	var uintptrType uintptr
	uintptrType = uintptr(unsafe.Pointer(&complexType))

	fmt.Printf("byteType = %d, intType = %d,  unsignedInt = %d, floatType = %.2f, complexType = %.2f, uintptrType = %d\n",
		byteType, intType, unsignedInt, floatType, complexType, uintptrType)
}

func nonNumericDataTypes() {
	var isTrue bool
	isTrue = false

	var myName string
	myName = "Nagakishore Sidde"

	// In between back quotes any character may appear except back quote.
	// Special characters (such as '\n' etc.) will not be interpreted,
	// in particular blackslashes have no special meaning
	rawString := `This is a raw string \n`
	fmt.Println(rawString)

	// In between double quotes any character may appear except unescaped double quotes.
	// Backslash escapes are interpreted
	interpretedString := "This is an \n\n interpreted string"
	fmt.Println(interpretedString)

	// define constant
	const Pi float64 = 3.14
	 //pi = 4.56 // Not allowed. Can not assign to constant

	// Other derived types such as struct, Array, slice, map etc. will be discussed individually

	utils.Header("Non-Numeric Data Types")
	fmt.Printf("isTrue = %t, name = %s, pi = %.2f\n", isTrue, myName, Pi)
}

func dataTypesFormatting() {
	age := 30
	pi := 3.14

	// To know the datatype, try the following
	fmt.Printf("Datatype of Age is '%T' \n", age)
	fmt.Printf("Binary representation of 5 : '%b' \n", 5)
	fmt.Printf("Character code for 65 is : '%c' \n", 65)
	fmt.Printf("Octal representation of 15 : '%o' \n", 15)
	fmt.Printf("Hexadecimal representation of 19 : '%x' \n", 19)
	fmt.Printf("Scientific notation of pi value is : '%e'\n", pi)
}

func arithmeticOperations() {
	utils.Header("Arithmetic Operations")
	fmt.Println("8 + 2 = ", 8+2)
	fmt.Println("8 - 2 = ", 8-2)
	fmt.Println("8 * 2 = ", 8*2)
	fmt.Println("8 / 2 = ", 8/2)
	fmt.Println("8 % 2 = ", 8%2)
}

func comparisionOperations() {
	utils.Header("Comparision Operations")
	fmt.Println("Is 2 == 2 ?", 2 == 2)
	fmt.Println("Is 2 != 2 ?", 2 != 2)
	fmt.Println("Is 2 <= 8 ?", 2 < 8)
	fmt.Println("Is 2 >= 8 ?", 2 > 8)

}
func logicalOperations() {
	utils.Header("Logical Operations")
	fmt.Println("true && false = ", true && false)
	fmt.Println("true || false = ", true || false)
	fmt.Println("!true = ", !true)
}

func keywords() {
	// Do not use these as Identifiers
	// var chan int // Not valid
	/*
		break        default      func         interface    select
		case         defer        go           map          struct
		chan         else         goto         package      switch
		const        fallthrough  if           range        type
		continue     for          import       return       var
	 */

}