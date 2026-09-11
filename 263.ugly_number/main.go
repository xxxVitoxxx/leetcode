package main

// time complexity: O(log n)
// space complexity: O(1)
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	for _, d := range []int{2, 3, 5} {
		for n%d == 0 {
			n /= d
		}
	}

	return n == 1
}
