package main

import "fmt"

func main() {
	sampleSet := make(Set)
	sampleSet["key1"] = struct{}{}
	sampleSet["key2"] = struct{}{}
	// There will be only one key2 in the set
	sampleSet["key2"] = struct{}{}

	fmt.Println("Set contents : ", getSetValues(sampleSet))
}

type Set map[string]struct{}

func getSetValues(s Set) []string {
	var resArr []string
	for key, _ := range s {
		resArr = append(resArr, key)
	}

	return resArr
}
