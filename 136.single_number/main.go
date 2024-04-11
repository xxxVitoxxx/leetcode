package main

// brute force
func singleNumber(nums []int) int {
	// -3 * 10^4 <= nums[i] <= 3 * 10^4
	appears := make([]int, 60001)
	for _, num := range nums {
		appears[num+30000]++
	}

	for i := range appears {
		if appears[i] == 1 {
			return i - 30000
		}
	}

	return -1
}

// use hash
// time complexity: O(n)
// space complexity: O(n)
func singleNumber2(nums []int) int {
	appears := make(map[int]int)
	for _, num := range nums {
		appears[num]++
	}

	for k, v := range appears {
		if v == 1 {
			return k
		}
	}

	return -1
}

// use bitwise operator
// time complexity: O(n)
// space complexity: O(1)
//
// 0 ^ 0 = 0
// 1 ^ 0 = 1
// 0 ^ 1 = 1
// if element appears twice, xor will be 0
func singleNumber3(nums []int) int {
	var xor int
	for _, num := range nums {
		xor ^= num
	}
	return xor
}
