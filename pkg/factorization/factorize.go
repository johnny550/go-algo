package factorization

import (
	"math"
)

// ref: https://learning.oreilly.com/videos/learning-data-structures/9781771373470/9781771373470-video209820/
// rules of the algo
// 1- don't check multiples of 2 after extracting them. (skipping even numbers)
// 2- don't check multiples of earlier values. (this rule is skipped in this implementation)
// 3- check values up to the square root of the current number
func Factorize(n int) []int {
	factors := []int{}
	factor := 2
	// get as many divisions by 2 as possible (even numbers)
	for n%factor == 0 {
		factors = append(factors, factor)
		n /= factor
	}

	// start from the next possible factor: 2+1
	factor = 3
	stopAt := math.Sqrt(float64(n))

	for factor <= int(stopAt) {
		for n%factor == 0 {
			factors = append(factors, factor)
			n /= factor
			stopAt = math.Sqrt(float64(n))
		}
		// skip even numbers because they are multiples of 2. (i+=2)
		factor += 2
	}
	// n is probably a prime number, add it to the factors
	if n > 1 {
		factors = append(factors, n)
	}

	return factors
}

// another implementation
// for i := 3; i <= n; i += 2 {
// 		fmt.Printf("stopat: %v  n: %v  factor: %v\n", stopAt, n, factor)

// 		for n%i == 0 {
// 			factors = append(factors, i)
// 			n /= i
// 		}
// 		if n > 1 {
// 			factors = append(factors, n)
// 			break
// 		}
// 	}
