package main

import (
	"fmt"
)

type Animal struct {
	Name string
	Age int
	Species string
	Weight float64
}

func PrintAnimal(who Animal) {
	fmt.Println("Your pet:", who)
}

func main() {
	frankie := Animal{"Frankie", 5, "Dog", 20.5}

	fmt.Println(frankie.Name)
}
