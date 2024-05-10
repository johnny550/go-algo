package main

import (
	"fmt"
	ds "go-algo/pkg/data_structures"
	"go-algo/pkg/factorization"
	"go-algo/pkg/sorting"
)

func main() {
	// trees
	n := ds.TreeNode{Key: 13}

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
	// fmt.Printf("stack: %v\n", stack)
	// stack.Push(2)
	// fmt.Printf("stack: %v\n", stack)
	// fmt.Printf("Peek: %v\n", stack.Peek())

	// fmt.Printf("Stack before Pop: %v\n", stack)
	// stack.Pop()
	// fmt.Printf("Stack after Pop: %+v\n", stack)

	// Sorting
	// InsertSort
	// sl := []int{2, 7, 8, 9}
	// sl = sorting.InsertSort(sl, 4)

	unsorted := []int{4, 9, 3, 7, 8, 2, 1}
	// sorting.SelectionSort(unsorted)
	sorting.QuickSort(unsorted)
	fmt.Printf("sorted: %v\n", unsorted)

	unsorted = []int{5, 6, 7, 2, 1, 0}
	sorting.QuickSort(unsorted)
	fmt.Printf("sorted: %v\n", unsorted)

	m := &sorting.MaxHeap{}
	fmt.Printf("Initial status of the heap: %v\n", m)
	buildHeap := []int{10, 30, 20, 75}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Printf("m: %v\n", m)
	}

	for i := 0; i < 3; i++ {
		m.Extract()
		fmt.Printf("---m: %v\n", m)
	}
}
