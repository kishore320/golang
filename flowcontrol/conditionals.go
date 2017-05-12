package main

import (
	"fmt"
	"runtime"
	"math/rand"
)

func main() {
	demoIf()
	demoSwitch()
	demoGoto()
	demoContinue()
}

func demoIf() {
	numCpus := runtime.NumCPU()

	// No need of parentheses around condition. Curly braces are mandatory
	 if numCpus := runtime.NumCPU(); numCpus > 1 { // in-line initialization can also be done like this
	// if numCpus > 1 {
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

func demoGoto() {
	for i := rand.Intn(10); ; i++ {
		for j := rand.Intn(10); ; j++ {
			fmt.Printf("i = %d, j = %d \n", i, j)
			if i > j {
				goto cleanup
				// goto innerBlock // Will not work
			} else {
				// innerBlock : {
				//	 fmt.Println("Do not use this block for goto")
				//	 return
				// }
				break
			}
		}
	}

	fmt.Println("I do not want this to be printed")

	cleanup : {
		fmt.Printf("Cleaning up before returning \n")
		return
	}
}

func demoContinue() {
	// Print even numbers only
	for i := 1; i < 10; i++ {
		if (i % 2 != 0) {
			continue
		}
		fmt.Printf("%d ", i)
	}
}
