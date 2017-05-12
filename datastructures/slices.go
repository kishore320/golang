package main

import "fmt"
import (
	"../utils"
)

func slices() {
	utils.Header("Slice :")

	// Slice type is an abstraction built on top of an array. Arrays are inflexible but slices are not.
	// Create a slice similar to an Array but leave out the size
	slice := []string{"hello", "world"}
	fmt.Println("Slice :", slice)

	// Built-in function 'make' can be used to create a slice
	// make([] T, initialLength, capacity) : make allocates an array and returns a slice that refers to an array
	// if capacity value is omitted it defaults to the specified length
	anotherSlice := make([] int, 4, 6)
	anotherSlice[0] = 2
	anotherSlice[3] = 8
	fmt.Println("New slice :", anotherSlice)
	fmt.Println("Length of the new slice :", len(anotherSlice))
	fmt.Println("Capacity of the new slice :", cap(anotherSlice))

	// Slices can be formed by "Slicing" an existing an array
	array := [4]int {2, 4, 6, 8}
	// Entire array will be sliced
	slicedFromArray := array[:]
	fmt.Println("Slice sliced from an array :", slicedFromArray)

	// Slices can be formed by "Slicing" an existing slice
	// Why does the following works ?
	slicedFromSlice := anotherSlice[2:6]
	fmt.Println("Slice sliced from a slice :", slicedFromSlice)

	/*
	    slicedFromSlice := anotherSlice[:4] // From index 0 to 3 will be sliced

	    slicedFromSlice := anotherSlice[1:] // From index 1 to till end of the array will be sliced

	    slicedFromSlice := anotherSlice[1:5] // From index 1 to 4 will be sliced
	*/

	// Two dimensional slices
	twoDSlice := [][]int { {1, 3}, {2, 4}}
	fmt.Println("Two dimensional Slice :", twoDSlice)
}

func sliceInternals() {
	utils.Header("Slice Internals")

	// A slice is a descriptor of an array segment. It consists of a pointer to an array,
	// length of the segment and its capacity. Slicing a slice does not copy the slice's data.
	// It creates a new slice value that points to the original array.

	slice := make([] int, 3, 6)
	slice[0] = 1
	slice[1] = 3
	slice[2] = 5
	fmt.Println("Original slice :", slice, "Length :", len(slice), "Capacity :", cap(slice))

	slicedSlice := slice[2:4]
	fmt.Println("Sliced slice :", slicedSlice, "Length :", len(slicedSlice), "Capacity :", cap(slicedSlice))
	/*
	 slice ->  | ptr | length (2) | capacity (6) |
	              |
	            _________________________
	            | 1 | 3 | 5 | 0 | 0 | 0 |
	                      |
	 slicedSlice ->    | ptr | length (2) | capacity (4) |
	*/

	slicedSlice[0] = 7
	fmt.Println("Original slice after modifying sliced slice :", slice, "Length :", len(slice), "Capacity :", cap(slice))
	fmt.Println("Sliced slice after modification  :", slicedSlice, "Length :", len(slicedSlice), "Capacity :", cap(slicedSlice))

	// Grow the slice length by slicing it again
	// slicedSlice : | ptr | length (4) | capacity (4) |
	slicedSlice = slicedSlice[:cap(slicedSlice)]
	fmt.Println("Sliced slice after growing :", slicedSlice, "Length :", len(slicedSlice), "Capacity :", cap(slicedSlice))

	// Slice can not grow beyond its capacity. The following will result in an error - runtime panic
	// slicedSlice = slicedSlice[:cap(slicedSlice) + 3]

	// Slices can not be re-sliced to below zero to access earlier elements in an array
	// slicedSlice = slicedSlice[-2:]  // This is not allowed
}

func growSliceCapacity() {
	utils.Header("Increase Slice Capacity")
	/*
		To increase slice's capacity one must create a new, larger size and copy the contents of original slice into it.
		Alternatively, one can use the built-in 'copy' function to achieve the same.
		Copy function supports copying between slices of different lengths and also handles source and destination
		slices that share the same underlying array.
	 */
	originalSlice := make([] int, 1) // Capacity is same as length
	originalSlice[0] = 2
	fmt.Println("Original Slice :", originalSlice, "Length :", len(originalSlice), "Capacity :", cap(originalSlice))

	copyOfOriginal := make([] int, 4, 8)
	copy(copyOfOriginal, originalSlice)
	fmt.Println("Copy of Original Slice :", copyOfOriginal, "Length :", len(copyOfOriginal),
		"Capacity :", cap(copyOfOriginal))

	// Change the copy and observe that original does not change
	copyOfOriginal[0] = 3
	fmt.Println("Original Slice after modifying the copy :", originalSlice, "Length :",
		len(originalSlice), "Capacity :", cap(originalSlice))

	// To achieve this one go using standard growing logic, use Append function. Observe how capacity changes
	originalSlice = append(originalSlice, 3)
	fmt.Println("Original Slice :", originalSlice, "Length :", len(originalSlice), "Capacity :", cap(originalSlice))

	originalSlice = append(originalSlice, 5)
	fmt.Println("Original Slice :", originalSlice, "Length :", len(originalSlice), "Capacity :", cap(originalSlice))

	originalSlice = append(originalSlice, 7, 9, 11)
	fmt.Println("Original Slice :", originalSlice, "Length :", len(originalSlice), "Capacity :", cap(originalSlice))

}

func sliceLenVsCapacity() {
	utils.Header("Slice: Length VS Capacity")
	slice := make([]int, 2, 4)
	for i := 1; i < 9; i++ {
		slice = append(slice, i)
		fmt.Println("Length :", len(slice), "Capacity :", cap(slice))
	}

	// What does it print ?
	fmt.Println("Why leading zeros ? ", slice)
}

func sliceTricks() {
	utils.Header("Slice Tricks :")

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Original slice : ", slice)

	// Remove elements from index 2 to 4
	slice = append(slice[:2], slice[5:]...)
	fmt.Println("Slice after removing elements from index 2 to 4 : ", slice)

	// Remove single element say at index 1
	slice = append(slice[:1], slice[2:]...)
	fmt.Println("Slice after removing element at index 1 : ", slice)
}

func main() {
	slices()
	//sliceInternals()
	//growSliceCapacity()
	//sliceTricks()
	//sliceLenVsCapacity()
}