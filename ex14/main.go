package main

import (
	"fmt"
)

func PrintThis(this string) {
	fmt.Println(this)
}

func main() {
	PrintThis("START")
	defer PrintThis("one")
	defer PrintThis("two")
	defer PrintThis("three")
	defer PrintThis("four")
	defer PrintThis("five")
	PrintThis("END")
}
