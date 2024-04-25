package main

import (
	"fmt"
	ds "go-algo/pkg/data_structures"
	"go-algo/pkg/factorization"
	"go-algo/pkg/trees"
)

func main() {
	// trees
	n := trees.Node{Key: 13}

	n.Insert(15)
	n.Insert(10)
	n.Insert(12)
	n.Insert(11)

	found := n.Search(11)
	// fmt.Printf("n: %v\n", n)
	fmt.Printf("found: %v\n", found)

	// Factors
	i := 5
	_ = factorization.Factorize(i)
	// fmt.Printf("factors for %d: %v\n", i, factors)

	// get prime numbers
	i = 10
	_ = factorization.GetPrime(i)
	// fmt.Printf("prime numbers under %d: %v\n", i, primeNumbers)

	// Stacks
	stack := ds.Stack[int]{}
	stack.Push(1)
	fmt.Printf("stack: %v\n", stack)
	stack.Push(2)
	fmt.Printf("stack: %v\n", stack)
	fmt.Printf("Peek: %v\n", stack.Peek())

	fmt.Printf("Stack before Pop: %v\n", stack)
	stack.Pop()
	fmt.Printf("Stack after Pop: %+v\n", stack)
}
