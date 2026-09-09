package main

import "strconv"

func isPalindrome(x int) bool {
	xStr := strconv.Itoa(x)
	for left, right := 0, len(xStr)-1; left < right; {
		if xStr[left] != xStr[right] {
			return false
		}

		left++
		right--
	}
	return true
}

// follow up: can not converting the integer to a string
// time complexity: O(log10(N)). we divided the x by 10 for every iteration.
// space complexity: O(1)
func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}

	var result int
	for curr := x; curr > 0; curr /= 10 {
		result = result*10 + curr%10
	}
	return result == x
}
