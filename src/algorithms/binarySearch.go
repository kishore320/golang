package main

import "fmt"

func main() {
	arr := [10]int{100, 89, 34, 55, 23, 11, 9, 45, 78, 999}

	// Should print true
	fmt.Println("Does 23 exists in the array ? ", binarySearchRecursive(arr[:], 23, 0, len(arr)-1))

	// Should print false
	fmt.Println("Does 231 exists in the array ? ", binarySearchRecursive(arr[:], 231, 0, len(arr)-1))

	// Should print true
	fmt.Println("Does 23 exists in the array ? ", binarySearchIterative(arr[:], 23))

	// Should print false
	fmt.Println("Does 231 exists in the array ? ", binarySearchIterative(arr[:], 231))
}

func binarySearchRecursive(arr []int, key int, low int, hi int) bool {
	if hi < low {
		return false
	}

	mid := (low + hi) / 2
	if arr[mid] == key {
		return true
	} else if key < arr[mid] {
		return binarySearchRecursive(arr, key, low, mid-1)
	} else {
		return binarySearchRecursive(arr, key, mid+1, hi)
	}
}

func binarySearchIterative(arr []int, key int) bool {
	low := 0
	hi := len(arr) - 1

	for hi >= low {
		mid := (low + hi) / 2
		if arr[mid] == key {
			return true
		} else if key < arr[mid] {
			hi = mid - 1
		} else {
			low = mid + 1
		}
	}

	return false
}
