package main

import "fmt"

// comparable Constraint is built interface - for comparing generic values using ==, !=

func Equal[T comparable](a T, b T) bool {
	return a == b
}
func main() {

	result := Equal[float32](1.5, 1.5)
	if result {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	result = Equal[int32](20, 20)
	if result {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	result = Equal[string]("Hello", "Hello")
	if result {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	result = Equal[bool](false, false)
	if result {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
