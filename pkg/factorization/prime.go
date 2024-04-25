package factorization

import (
	"fmt"
	"math"
)

// rules of the algo
// 1- don't check multiples of 2 after extracting them. (skipping even numbers)\
// 2- check values up to the square root of the current number
func GetPrime(n int) []int {
	// nonPrimeNumbersUnderI := make(map[int]bool)
	// nonPrimeNumbersUnderI[1] = false
	// factor := 2
	// // for n%factor == 0 {
	// // 	// primeNumbersUnderI = append(primeNumbersUnderI, factor)
	// // 	primeNumbersUnderI[factor] = true
	// // 	n /= factor
	// // }

	// // factor = 3
	// stopAt := math.Sqrt(float64(n))

	// for factor <= int(stopAt) {
	// 	fmt.Printf("---n: %d factor: %v  stopat: %v\n", n, factor, stopAt)
	// 	multiple := 2 * factor
	// 	for multiple <= n {
	// 		fmt.Printf("n: %d factor: %v  stopat: %v\n", n, factor, stopAt)
	// 		fmt.Printf("multiple: %v\n", multiple)
	// 		nonPrimeNumbersUnderI[multiple] = false
	// 		factor++
	// 		multiple = 2 * factor
	// 	}
	// 	factor += 2
	// 	stopAt = math.Sqrt(float64(n))
	// }
	// fmt.Printf("primeNumbersUnderI: %v\n", nonPrimeNumbersUnderI)
	// rst := []int{}
	// for i := 2; i <= n; i++ {
	// 	// if the current number i isn't in the nonPrime map, it means it is a prime number, return it
	// 	if _, ok := nonPrimeNumbersUnderI[i]; !ok {
	// 		rst = append(rst, i)
	// 	}
	// }
	// return rst

	switch {
	case n < 2:
		return []int{}
	case n == 2:
		return []int{2}
	}
	// a[i] == false ==> p=2*i+3 is a candidate prime
	// p in [3,n] ==> i in [0,(n-3)/2]
	// Formulae to know how many prime numbers <= current number: 1+ (n-3)/2
	length := 1 + (n-3)/2
	a := make([]bool, length)
	// Start with number 3 and consider only odd numbers
	sqrtn := int(math.Sqrt(float64(n)))
	for i, p := 0, 3; p <= sqrtn; p += 2 {
		fmt.Printf("i: %d  a[i]: %v\n", i, a[i])
		if !a[i] {
			// 2*i+1 is a prime number; mark off its multiples; step to skip=the prime candidate (2 in this example)
			for j := (p*p - 3) / 2; j < length; j += p {
				a[j] = true
			}
		}
		i++
	}
	// ps will store the computed primes;
	ps := make([]int, 1)
	ps[0] = 2
	for i := 0; i < length; i++ {
		if !a[i] {
			ps = append(ps, 2*i+3)
		}
	}
	return ps
}
