package main

import "fmt"

func printValue[T any](value T) {
	fmt.Println(value)
}

func main() {

	printValue(10)
	printValue(12.5)
	printValue("Type Constraint Example")
	printValue(false)
}
