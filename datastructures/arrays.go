package main

import (
	"fmt"
	"github.com/kishore320/golang/utils"
)

func arrayBasics() {
	utils.Header("Array Basics")

	// An array is a numbered sequence of elements which are of same type and its size is fixed
	// Declare an array of ints whose size is 2
	var evenNumbers [2]int
	// Must be initialized automatically with default values. In this case 0's
	fmt.Println("Even numbers before explicit assignment :", evenNumbers)

	// Now start assigning
	evenNumbers[0] = 2
	evenNumbers[1] = 4
	// The following will error out with msg "invalid array index 2 (out of bounds for 2-element array)"
	// evenNumbers[2] = 6

	// Initialize and assign in a single line
	oddNumbers := [4]int {1, 3, 5, 7}
	fmt.Println("Odd numbers :", oddNumbers)

	// You can also let the compiler count the array elements for you using '...'
	primeNumbers := [...]int {2, 3, 5, 7}
	// Use function 'len' to get the array length
	fmt.Println("Length of an array :", len(primeNumbers))

	// Two dimensional array
	twoDArr := [2][2]int {
		{1, 3},
		{2, 4},
	}

	utils.Header("Two Dimensional Array")
	for i, iVal := range twoDArr {
		for j, jVal := range iVal {
			fmt.Printf("arr[%d][%d] = %d\n", i, j, jVal)
		}
	}

	// (OR) simply
	//fmt.Println(twoDArr)
}

func demoOfDifference() {
	utils.Header("Some Differences when compared to 'C'")

	// Go's arrays are values and not pointers to the first element of an array (unlike in 'C')
	// An array variable denotes an entire array, so assigning array variable (or) passing it around
	// makes a copy of its contents
	evenNumbers := [2]int {2, 4}
	fmt.Println("Original Array :", evenNumbers)

	differentCopyWithAssignment := evenNumbers
	// Change the copy and observe that original array does not change
	differentCopyWithAssignment[0] = 3
	differentCopyWithAssignment[1] = 5

	fmt.Println("Original Array after modifying the copied array :", evenNumbers)
	fmt.Println("Copied Array :", differentCopyWithAssignment)

	// Same is true when array variable passed around. A brand new copy will be made
	convertEvenToOdds(evenNumbers)
	fmt.Println("After converting evens to odds with out using pointers :", evenNumbers)

	// To avoid the copy, use pointers
	convertEvenToOddsUsingPointer(&evenNumbers)
	fmt.Println("After converting evens to odds using pointers :", evenNumbers)
}

func convertEvenToOdds(arr [2]int) {
	arr[0] = 3
	arr[1] = 5
}

func convertEvenToOddsUsingPointer(arr *[2]int) {
	arr[0] = 3
	arr[1] = 5
}

func main() {
	//arrayBasics()
	demoOfDifference() // when compared to other languages say 'C', 'Java'
}
