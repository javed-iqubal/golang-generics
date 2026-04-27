package main

import "fmt"

// custom constraint
type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Add[T Number](a T, b T) T {
	return a + b
}
func main() {

	fmt.Println(Add(10, 15.5))
	fmt.Println(Add(10.25, 15.5))
	fmt.Println(Add(10.35, 45))

}
