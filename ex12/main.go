package main

import "fmt"

func main() {
	fruit := "Apple"

	switch fruit {
	case "Apple":
		fmt.Println("You have Apples.")
	case "Orange":
		fmt.Println("You have Oranges.")
	case "Grape":
		fmt.Println("You have Grapes.")
	case "Durian":
		fmt.Println("You have Durian.")
	default:
		fmt.Println("You have some fruit.")
	}

	// You can also use a switch like a "super if"
	i := 23

	switch {
	case i < 20:
		fmt.Println("I is less than 20")
	case i > 20:
		fmt.Println("I is greater than 20")
	default:
		fmt.Println("Then what is i now?")
	}
}
