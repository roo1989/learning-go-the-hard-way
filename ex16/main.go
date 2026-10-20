package main

import (
	"fmt"
)

func main() {
	var age int
	var name string
	var favorite string

	fmt.Println("What's your name? ")
	fmt.Scan(&name)

	fmt.Println("How old are you? ")
	fmt.Scanln(&age)

	fmt.Println("What's your favorite kind of pet? ")
	fmt.Scan(&favorite)

	fmt.Println("Your name is", name)
	fmt.Println("You are", age, "years old")
	fmt.Println("Your favorite pet is", favorite)
}
