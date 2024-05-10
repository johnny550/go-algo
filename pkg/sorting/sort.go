package sorting

import "fmt"

// for each item:
// - find the position where the item belongs
// - insert the item
func InsertSort(sl []int, numToInsert int) []int {
	for i := 0; i <= len(sl)-1; i++ {
		if sl[i] < numToInsert {
			continue
		}

		// temp will hold the part of sl to move forward when the right spot to insert NUM is found
		temp := []int{}
		temp = append(temp, sl[i:]...)
		sl[i] = numToInsert
		sl = append(sl[:i+1], temp...)
		break
	}
	return sl
}

// for p=1 to N
// - find the smallest remaining item, say in position k
// - swap the items in positions p & k
// 2 - 9 - 3 - 7  [step 0]
// |____|
// 2 - 9 - 3 - 7  [step 1]
//
//	|____|
//
// 2 - 3 - 9 - 7  [step 2]
//
//	|____|
//
// 2 - 3 - 7 - 9  [step 3]
func SelectionSort(sl []int) {
	left, right := 0, 1
	for left = 0; left <= len(sl)-1; left++ {
		fmt.Printf("sl at step[%d]: %v\n", left, sl)
		if right == len(sl) {
			break
		}
		if sl[left] > sl[right] {
			sl[left], sl[right] = sl[right], sl[left]
		}
		right++
	}
}

// for each in array
// - pick a divider
// - divide into 2 groups. group 1 < divider. group 2 > divider
// - recursively call Quicksort
// O(NlogN): if elements are non sorted originally
// O(N)^2: if elements are originally sorted or come with many duplicates
func QuickSort(sl []int) {
	if len(sl) <= 1 {
		return
	}
	divider := sl[0]
	// materializing the hole. How since it's slice of int? can't have nil. for simplicity, let's use 0
	// sl[0] = 0
	// hole on the left side of the slice. so start by evaluating the value at the right most of the slice
	hole := 0
	right := len(sl) - 1
	left := 1
	// fmt.Printf("\n\nbefore: %v\n", sl)
	i := 0

	for left <= right {
		if i%2 == 0 {
			// eval right
			fmt.Println("Evaluating the right")
			for sl[right] > divider {
				// that number is where it belongs, the right of the slice
				// move the index to the left and eval the right side again
				// whenever val @right is bigger|eq to divider
				right--
				continue
			}
			if sl[right] < divider {
				// to the left
				// fmt.Printf("Evaluating %d against %d\n", sl[right], sl[hole])
				sl[hole], sl[right] = sl[right], sl[hole]
				// new emplacement of the hole (old emplacement of right)
				hole = right
			}
			// fmt.Printf("left: %d  right:%d\n\n", left, right)
			right--

			// fmt.Printf("sl: %v\n", sl)
			i++
			continue
		}
		// eval left
		for sl[left] < divider {
			// that number is where it belongs, the left of the slice
			// move the index to the right and eval the left side again
			// whenever val @right is bigger|eq to divider
			left++
			continue
		}
		if sl[left] > divider {
			// fmt.Println("Evaluating the left")
			// fmt.Printf("--Evaluating %d against %d\n", sl[hole], sl[left])
			sl[hole], sl[left] = sl[left], sl[hole]
			// new emplacement of the hole (old emplacement of left)
			hole = left
			// fmt.Printf("--left: %d  right:%d\n\n", left, right)
		}
		left++
		// fmt.Printf("--sl: %v\n", sl)
		i++
	}
	leftSide := sl[:hole]
	rightSide := sl[hole+1:]
	QuickSort(leftSide)
	QuickSort(rightSide)
	// fmt.Printf("leftSide: %v\n", QuickSort(leftSide))
	// fmt.Printf("rightSide: %v\n", rightSide)
}
