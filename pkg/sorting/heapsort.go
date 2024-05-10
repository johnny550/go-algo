package sorting

// maxHeap contains a slice holding the array of elements
type MaxHeap struct {
	heapArr []int
}

// Adds an element to the heap
func (m *MaxHeap) Insert(key int) {
	// add to the bottom of the arr (bottom left in tree form)
	m.heapArr = append(m.heapArr, key)
	// re-arrange the inserted key and validate the heap
	m.maxHeapifyUp(len(m.heapArr) - 1)
}

// Extracts the largest key and removes it from the heap
func (m *MaxHeap) Extract() int {
	if len(m.heapArr) == 0 {
		return -1
	}

	largestKey := m.heapArr[0]

	l := len(m.heapArr) - 1
	// take the last el in the heap, make of it the root
	m.heapArr[0] = m.heapArr[l]
	// shorten the array. up to but excluding the last element
	m.heapArr = m.heapArr[:l]
	// heapify but this time, from the top to the bottom
	m.maxHeapifyDow(0)

	return largestKey
}

func (m *MaxHeap) maxHeapifyDow(index int) {
	leftChild, rightChild := getLeftChild(index), getRightChild(index)
	lastIndex := len(m.heapArr) - 1
	childToCompare := 0

	// loop while el @index has at least 1 child
	for leftChild <= lastIndex {
		// when leftChild is the only child
		if leftChild == lastIndex {
			childToCompare = leftChild
		} else if m.heapArr[leftChild] > m.heapArr[rightChild] {
			// when lc is larger than rc
			childToCompare = leftChild
		} else {
			// when rc is larger
			childToCompare = rightChild
		}

		// compare the arr value @index to the largest child, swap'em if needed
		if m.heapArr[index] < m.heapArr[childToCompare] {
			m.swap(index, childToCompare)
			index = childToCompare
			leftChild, rightChild = getLeftChild(index), getRightChild(index)
		}
		// otherwise, the element is at the right spot. stop the heapify
		return
	}
}

// heapify from bottom to top: place an element at the bottom of the heap then find its right place,
// swapping when necessary, going up the heap
func (m *MaxHeap) maxHeapifyUp(index int) {
	// as long as the curr el at index is bigger than its parent, swap the 2
	for m.heapArr[getParent(index)] < m.heapArr[index] {
		m.swap(getParent(index), index)
		index = getParent(index)
	}
}

// Gets the index of the parent node
func getParent(i int) int {
	return (i - 1) / 2
}

// Gets the index of the left child
func getLeftChild(i int) int {
	return (i * 2) + 1
}

// Gets the index of the right child
func getRightChild(i int) int {
	return (i * 2) + 2
}

// swaps the value of 2 indices
func (m *MaxHeap) swap(i, j int) {
	m.heapArr[i], m.heapArr[j] = m.heapArr[j], m.heapArr[i]
}
