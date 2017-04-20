package main

import "fmt"
import "../utils"

func main() {
	arrayDemo()
	sliceDemo()
	sliceTricks()
	sliceOperations()
	sliceLenVsCapacity()
}

func arrayDemo() {
	utils.Header("Array :")

	// Array is fixed size. We can not grow it.
	var array [2]string
	array[0] = "hello"
	array[1] = "world"
	fmt.Println(array)

	// Another way to create an array
	arrayAnother := [2]string{"hello", "world"}
	for index, value := range arrayAnother {
		fmt.Println(index, value)
	}
}

func sliceDemo() {
	utils.Header("Slice :")

	slice := []string{"hello", "world"}
	// Unlike arrays, slice can grow dynamically
	slice = append(slice, "how", "are", "you")
	fmt.Println(slice)

	// Create a slice with length 2 and capacity 5
	anotherSlice := make([]int, 2, 5)
	// should print just two 0s which is a length of the slice
	fmt.Println(anotherSlice)

	// initialize values
	anotherSlice[0] = 1
	anotherSlice[1] = 2
	anotherSlice = append(anotherSlice, 3, 4, 5)
	fmt.Println(anotherSlice)
}

func sliceOperations() {
	utils.Header("Slice Operations :")

	// Unlike arrays, while defining a slice, leave out the size
	slice := []int{3, 4, 12, 8, 90, 56, 77, 100, 34}
	slice1 := slice[4:6] // should have from index 4 to 5 (excluding 6)
	slice2 := slice[:4]  // should have from starting index (i.e 0) to 3 (excluding 4)
	slice3 := slice[6:]  // should have from index 6 to all the way end

	fmt.Println("Original slice : ", slice)
	fmt.Println("Slice1 slice[4:6] : ", slice1)
	fmt.Println("Slice2 slice[:4] : ", slice2)
	fmt.Println("Slice3 slice[6:] : ", slice3)

	// Interesting. Note that sub slices will point to the original slice
	// So, changes to the sub slices will change the original slice
	slice3[0] = 999
	fmt.Println("Slice3 after changing slice3[0] value to 999 : ", slice3)
	fmt.Println("Original slice after changing slice3[0] value to 999 : ", slice)

	// Append new elements to the existing slice
	slice = append(slice, 200, 345, 500)
	fmt.Println("Slice after appending 200, 345, 500 : ", slice)

	// Another way to create a slice which just allocates memory
	a := make([]int, 5, 10) // Length of the slice is 5 and the capacity of the slice is 10
	for i := 0; i < 5; i++ {
		a[i] = i + 1
	}

	fmt.Println("Slice values after just allocating memory : ", a)
	fmt.Println("Length of the slice : ", len(a))
	fmt.Println("Capacity of the slice : ", cap(a))

	// Slice will just change the pointer and will have access to the rest of the original slice also
	aSlice := a[2:6]
	fmt.Println("Slice contents : ", aSlice)
	fmt.Println("Slice contents using capacity : ", aSlice[:cap(aSlice)])

	// Use copy to create a brand new slice instead of just changing the pointer
	newSlice := make([]int, 3)
	copy(newSlice, a[2:5])
	fmt.Println("New copy of slice contents : ", newSlice[:cap(newSlice)])

	newSlice = append(newSlice, 6, 7, 8, 9)
	fmt.Println("Contents of slice after append : ", newSlice)
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
