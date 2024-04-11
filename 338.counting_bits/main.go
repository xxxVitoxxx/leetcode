package main

// pop count or hamming weight
// time complexity: O(n * logn)
// space complexity: O(1)
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
// space complexity: O(n)
func countBits3(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i < len(ans); i++ {
		ans[i] = ans[i>>1] + i&1
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
