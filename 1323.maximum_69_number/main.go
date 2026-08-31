package main

import "math"

// time complexity: O(1)
// the problem constraint num <= 10**4 means the loop runs at most 4 times.
// since the number of operations is capped by a small constant, the runtime is O(1).
//
// space complexity: O(1)
// we only use a few variables for our calculations.
// the space required does not scale with the input.
func maximum69Number(num int) int {
	dights := int(math.Log10(float64(num)))
	n := num
	for i := dights; i >= 0; i-- {
		powerOf10 := int(math.Pow(10, float64(i)))
		quotient := n / powerOf10
		if quotient == 6 {
			return num + 3*powerOf10
		}

		n %= quotient * powerOf10
	}

	return num
}
