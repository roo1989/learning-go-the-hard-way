package main

func WontWork(a int) {
	a := 1
}

func main() {
	a := 100
	WontWork(a)
}
