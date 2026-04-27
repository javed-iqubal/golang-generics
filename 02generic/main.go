package main

import "fmt"

// generic function

// func FunctionName[placeHolder Type](args Type) Type{}
func add[T int | float32 | float64](a, b T) T {
	return a + b
}

func main() {

	fmt.Printf("Addition of int is : %d \n", add(10, 20))
	fmt.Printf("Addition of float is : %f\n", add(10.25, 20.50))
	// error because string in not in generic type
	// fmt.Printf("Addition of string is : %s\n", add("Hello", "world"))
}
