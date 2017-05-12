package main

import "fmt"

func main() {
	pf := partialFunction(5)
	fmt.Println(pf(5))
}

func partialFunction(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
