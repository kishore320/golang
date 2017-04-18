package utility

import "fmt"

func SayHello() {
	fmt.Println("Hello World from utility")
}

// This will not be exported, so cant be used from other files
func sayNewHello() {
	fmt.Println("Non visible function")
}
