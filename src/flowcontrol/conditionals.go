package main

import (
	"fmt"
	"runtime"
)

func main() {
	demoIf()
	demoSwitch()
}

func demoIf() {
	numCpus := runtime.NumCPU()

	// No need of parentheses around condition. Curly braces are mandatory
	if numCpus > 1 {
		fmt.Println("Your system is a multi CPU system")
	} else {
		fmt.Println("Your system is a single CPU system")
	}
}

func demoSwitch() {
	// Observe that there is no 'break' statement
	switch runtime.GOOS {
	case "darwin":
		fmt.Println("You are using MAC OS")
	case "linux":
		fmt.Println("You are using Linux OS")
	case "freebsd":
		fmt.Println("You are using Free BSD OS")
	default:
		fmt.Println("Which OS are you using ?")
	}
}
