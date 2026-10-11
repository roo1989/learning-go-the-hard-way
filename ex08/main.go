package main

import (
	"fmt"
)

func main() {
	sunshine := true
	rain := false

	if sunshine && !rain {
		fmt.Println("I want to go outside to paint.")
	} else if sunshine && rain {
		fmt.Println("That's nice but I can't paint.")
	} else if !sunshine && rain {
		fmt.Println("It's really raining out.")
	} else if !sunshine && !rain {
		fmt.Println("I can go out, but there is no light.")
	} else {
		panic("The world ends.")
	}
}
