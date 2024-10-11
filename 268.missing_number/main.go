package main

import (
	"slices"
)

// use hash
// time complexity: O(n)
// space complexity: O(n)
func missingNumber(nums []int) int {
	m := make(map[int]bool, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		m[i] = false
	}

	for _, n := range nums {
		if _, ok := m[n]; ok {
			m[n] = true
		}
	}

	for k, v := range m {
		if !v {
			return k
		}
	}

	return -1
}

// use slice
// time complexity: O(n)
// space complexity: O(n)
func missingNumber2(nums []int) int {
	arr := make([]bool, len(nums)+1)
	for _, n := range nums {
		arr[n] = true
	}

	for i := 0; i < len(arr); i++ {
		if !arr[i] {
			return i
		}
	}

	return -1
}

// use sort
// time complexity: O(nk) an O(nk) worst case of pdqsort with k distinct elements
// space complexity: O(1)
func missingNumber3(nums []int) int {
	slices.Sort(nums)

	if nums[0] != 0 {
		return 0
	}
	if nums[len(nums)-1] != len(nums) {
		return len(nums)
	}

	for i, n := range nums {
		if i != n {
			return i
		}
	}

	return -1
}

// use XOR
// time complexity: O(n)
// space complexity: O(1)
func missingNumber4(nums []int) int {
	missing := len(nums)
	for i, n := range nums {
		missing ^= i ^ n
	}
	return missing
}
