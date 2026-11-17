package main

import "math"

// 32-bit
// the max value is 2147483647
// the min value is -2147483648
// time complexity: O(log(x))
// space complexity: O(1)
func reverse(x int) int {
	var num int
	for x != 0 {
		pop := x % 10
		x /= 10
		if num > math.MaxInt32/10 || (num == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		if num < math.MinInt32/10 || (num == math.MinInt32/10 && pop < -8) {
			return 0
		}

		num = num*10 + pop
	}
	return num
}
