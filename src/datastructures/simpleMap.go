package main

import "fmt"

func main() {
	// One way to create map of string, string
	var kvMap map[string]string
	kvMap = make(map[string]string)
	kvMap["1"] = "one"
	kvMap["2"] = "two"

	// Another way
	kvMapNew := map[string]string{
		"3": "three",
		"4": "four",
	}

	fmt.Println("Value of key '1' :", kvMap["1"])
	fmt.Println("Value of key '3' :", kvMapNew["3"])
	fmt.Println("Map size is :", len(kvMapNew))

	// Check whether key exists in the map or not
	val, exists := kvMap["2"]
	fmt.Println(val)
	fmt.Println(exists)

	// Delete from the map
	delete(kvMap, "1")
	_, exists1 := kvMap["1"] // exists1 must be false
	fmt.Println("Does '1' exist in the map ?", exists1)

	for key, value := range kvMapNew {
		fmt.Println(key, ": ", value)
	}
}
