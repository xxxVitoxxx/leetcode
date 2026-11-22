package main

// time complexity: O(N)
// space complexity: O(N)
func plusOne(digits []int) []int {
	result := make([]int, len(digits)+1)
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		num := digits[i]
		num += carry

		result[i+1] = num % 10
		carry = num / 10
	}

	if carry == 0 {
		return result[1:]
	}

	result[0] = carry
	return result
}

// time complexity: O(N)
// space complexity: O(1)
func plusOne2(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}
