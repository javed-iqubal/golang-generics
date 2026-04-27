package main

import "fmt"

func Filter[T any](items []T, predicate func(T) bool) []T {

	var result []T

	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}
func main() {

	nums := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	predicateEven := func(n int32) bool {
		return n%2 == 0
	}

	evenNumbers := Filter(nums, predicateEven)

	fmt.Println("Even Numbers: ", evenNumbers)

	predicateOdd := func(n int32) bool {
		return n%2 != 0
	}

	oddNumbers := Filter(nums, predicateOdd)

	fmt.Println("Odd Numbers: ", oddNumbers)
}
