package main

// bit manipulation
// time complexity: O(1)
// space complexity: O(1)
//
// n = 15. the n binary string is 1111
// 1111 & 1 = 1. n >> 1 = 111
// 111 & 1 = 1. n >> 1 = 11
// 11 & 1 = 1. n >> 1 = 1
// 1 & 1 = 1. n >> 1 = 0
//
// n = 10. the n binary string is 1010
// 1010 & 1 = 0. n >> 1 = 101
// 101 & 1 = 1. n >> 1 = 10
// 10 & 1 = 0. n >> 1 = 1
// 1 & 1 = 1. n >> 1 = 0
func hammingWeight(n int) int {
	var result int
	for n != 0 {
		if n&1 == 1 {
			result++
		}
		n >>= 1
	}

	return result
}

// bitwise manipulation
// time complexity: O(1)
// space complexity: O(1)
//
// n = 15. the n binary string is 1111
// 1111 & 1110 = 1110. n = 1110
// 1110 & 1101 = 1100. n = 1100
// 1100 & 1011 = 1000. n = 1000
// 1000 & 111 = 0. n = 0
//
// n = 10. the n binary string is 1010
// 1010 & 1001 = 1000. n = 1000
// 1000 & 111 = 0. n = 0
func hammingWeight2(n int) int {
	var result int
	for n != 0 {
		n &= (n - 1)
		result++
	}
	return result
}
