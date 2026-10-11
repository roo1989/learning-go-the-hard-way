package main

import "fmt"

func main() {
	variables()
	constants()
	iterations()
}

func variables() {
	var age int = 37
	name := "Ragnar"

	fmt.Println("My name is", name, "and I am", age, "years old")

	var hair string
	hair = "Blond"

	fmt.Println("I have a", hair, "hairstyle.")

	var occupation string
	fmt.Println("Zero value of a string is:", occupation)

}

func constants() {
	const n = 5000000

	const d = 3e20 / n

	fmt.Println(d)

	fmt.Println(int64(d))
}

func iterations() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}
}
