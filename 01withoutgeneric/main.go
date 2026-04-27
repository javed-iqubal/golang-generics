package main

import "fmt"

func addInt(a, b int) int {
	return a + b
}

func addFloat64(a, b float64) float64 {
	return a + b
}

func main() {

	fmt.Printf("Addition of int is : %d \n", addInt(10, 20))
	fmt.Printf("Addition of float is : %f\n", addFloat64(10.25, 20.50))
}
