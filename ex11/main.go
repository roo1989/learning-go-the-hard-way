package main

import (
	"fmt"
)

func main() {
	versionTwo()
	practice()
	i := 0
loop:
	fmt.Println("loop!", i)
	i++

	if i < 10 {
		goto loop
	}
	fmt.Println("Loop ended")
}

func versionTwo() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("loope", i, "times")
	}
}

func practice() {
	// first style loos like an if-statement
	i := 10
	for i < 10 {
		fmt.Println("loop", i)
		i++
	}
	fmt.Println("------ first loop ended")

	// second style does all of that for you.
	for i := 0; i < 10; i++ {
		fmt.Println("loop", i)
	}

	fmt.Println("------- second loop ended")
}
