package main

import (
	"fmt"
)

func ReadName() string {
	fmt.Print("What is your name? ")

	var name string
	_, err := fmt.Scan(&name)
	if err != nil { panic("WRONG!") }

	return name
}

func main() {
	name := ReadName()
	fmt.Println("Thank you for your service citizen", name)
}
