package datastructures

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Stacks, with a dash of generics
type customData interface {
	constraints.Ordered | string | int
}

type Stack[T customData] struct {
	top    *Node[T]
	length int
}

type Node[T customData] struct {
	value    T
	previous *Node[T]
}

// type MyVal interface {
// 	int | float64
// }

func (s *Stack[T]) Push(i T) {
	// One way
	// newNode := &Node{value: i}
	// if s.length == 0 {
	// 	newNode.previous = nil
	// } else {
	// 	newNode.previous = s.top
	// }
	// Anotyher way. s.top will be nil if s.length=0
	newNode := &Node[T]{value: i, previous: s.top}

	s.top = newNode
	s.length++

	// DEBUG
	// fmt.Printf("s.top: %+v\n", s.top)
	// if s.top.previous != nil {
	// 	fmt.Printf("s.top.previous: %+v\n", s.top.previous)
	// }
}

func (s *Stack[T]) Peek() T {
	if s.length == 0 {
		return T(rune(0))
	}
	return s.top.value
}

func (s *Stack[T]) Pop() *Stack[T] {
	if s.length == 0 {
		return nil
	}
	fmt.Printf("s.top.previous: %v\n", s.top.previous)
	fmt.Printf("s.top: %v\n", s.top)
	// one way
	s.top = s.top.previous

	// another way
	// n := s.top
	// s.top = n.previous

	s.length--

	return s
}
