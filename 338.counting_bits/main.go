package main

// pop count or hamming weight
// time complexity: O(n * logn)
// space complexity: O(1)
//
// solve the problem for one number at time
// see leetcode 191
func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i < len(ans); i++ {
		ans[i] = popCount(i)
	}
	return ans
}

func popCount(n int) int {
	var result int
	for n != 0 {
		n &= (n - 1)
		result++
	}
	return result
}

// dynamic programming + most significant bit
// time complexity: O(n)
// space complexity: O(1)
func countBits2(n int) []int {
	ans := make([]int, n+1)
	x, b := 0, 1
	for b <= n {
		for x < b && x+b <= n {
			ans[x+b] = ans[x] + 1
			x += 1
		}
		x = 0
		b <<= 1
	}
	return ans
}

// dynamic programming + least significant bit
// time complexity: O(n)
// space complexity: O(1)
//
// ans[n] = ans[n/2] + n & 1
// binary  n  1's
//
//	  0    0   0
//	  1    1   1 ans[1] = ans[1/2] + 1 & 1 -> 1
//	 10    2   1 ans[2] = ans[2/2] + 2 & 1 -> 1
//	 11    3   2 ans[3] = ans[3/2] + 3 & 1 -> 2
//	100    4   1 ans[4] = ans[4/2] + 4 & 1 -> 1
//	101    5   2 ans[5] = ans[5/2] + 5 & 1 -> 2
//	110    6   2 ans[6] = ans[6/2] + 6 & 1 -> 2
//	111    7   3 ans[7] = ans[7/2] + 7 & 1 -> 3
func countBits3(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i < len(ans); i++ {
		ans[i] = ans[i/2] + i&1
		// ans[i] = ans[i>>1] + i&1
	}
	return ans
}

// dynamic programming + last set bit
// time complexity: O(n)
// space complexity: O(1)
func countBits4(n int) []int {
	ans := make([]int, n+1)
	for i := 1; i < len(ans); i++ {
		ans[i] = ans[i&(i-1)] + 1
	}
	return ans
}
