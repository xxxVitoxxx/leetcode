package main

// prefix sum
// time complexity: O(n)
// space complexity: O(1)
func largestAltitude(gain []int) int {
	num, ans := 0, 0
	for i := 0; i < len(gain); i++ {
		num += gain[i]
		if num > ans {
			ans = num
		}
	}
	return ans
}
