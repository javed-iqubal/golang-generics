package main

import "fmt"

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {

	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {

	if len(s.items) == 0 {
		panic("Stack is empty!")
	}
	n := len(s.items)
	item := s.items[n-1]
	s.items = s.items[:n-1]
	return item
}

func main() {
	stack := Stack[int]{}

	stack.Push(5)
	stack.Push(8)
	stack.Push(9)

	fmt.Println(stack.Pop())

	stack1 := Stack[string]{}

	stack1.Push("A")
	stack1.Push("C")
	stack1.Push("H")

	fmt.Println(stack1.Pop())
}
