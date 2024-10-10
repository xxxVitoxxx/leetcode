package main

// time complexity: O(1)
// space complexity: O(1)
func getSum(a, b int) int {
	for b != 0 {
		a, b = a^b, a&b<<1
	}
	return a
}

// time complexity: O(1)
// space complexity: O(1)
func getSum2(a, b int) int {
	if b == 0 {
		return a
	}
	return getSum2(a^b, a&b<<1)
}
