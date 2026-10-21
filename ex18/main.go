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
	PrintAnimal(frankie)

	frankie.Weight = 25.6
	fmt.Println("Your Pet Gained Weight:", frankie)

	miss_snuffles := frankie

	if miss_snuffles == frankie {
		miss_snuffles.Name = "Miss Snuffles"
		miss_snuffles.Species = "Cat"
		miss_snuffles.Weight = 4.0
	}

	PrintAnimal(frankie)
	PrintAnimal(miss_snuffles)
}
