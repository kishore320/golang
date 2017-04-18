package main

import "fmt"

func main() {
	forDemo()
	whileDemo()
}

func forDemo() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}

	fmt.Println("Sum from 1 to 100 is : ", sum)
}

func whileDemo() {
	// There is no 'while' keyword, use 'for' itself as follows..
	sum := 0
	i := 1
	for i <= 100 {
		sum += i
		i++
	}

	fmt.Println("Sum from 1 to 100 is : ", sum)
}
