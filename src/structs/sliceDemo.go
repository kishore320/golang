package main

import "fmt"

func main() {
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

	// Appropriate way to create slice and allocate memory
	a := make([]int, 10)
	fmt.Println("Slice values after just allocating memory : ", a)
}
