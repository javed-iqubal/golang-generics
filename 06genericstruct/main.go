package main

import "fmt"

type Box[T any] struct {
	value T
}

func (b Box[T]) Get() T {
	return b.value
}

func (b *Box[T]) Set(newValue T) {
	b.value = newValue
}

func main() {

	// Using box as integer
	intBox := Box[int]{100}
	intBox.Set(200)

	fmt.Println(intBox.Get())

	// Using box as integer
	float32Box := Box[float32]{1000.25}
	float32Box.Set(250.55)

	fmt.Println(float32Box.Get())
}
