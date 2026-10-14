package main

import (
	"fmt"
)

func Add(a float64, b float64) float64 {
	return a + b
}

func Sub(a float64, b float64) float64 {
	return a - b
}

func Mul(a float64, b float64) float64 {
	return a * b
}

func Div(a float64, b float64) float64 {
	return a / b
}

func main() {
	result := Add(Sub(3, 4), Mul(10, Div(200, 10)))

	fmt.Println("The result is:", result)
}
