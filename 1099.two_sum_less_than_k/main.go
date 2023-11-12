package main

import "sort"

// brute force
func twoSumLessThanK(nums []int, k int) int {
	res := -1
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			num := nums[i] + nums[j]
			if num > res && num < k {
				res = num
			}
		}
	}
	return res
}

// use two pointer
func twoSumLessThanK2(nums []int, k int) int {
	sort.Ints(nums)
	left, right, res := 0, len(nums)-1, -1
	for left < right {
		num := nums[left] + nums[right]
		if num < k {
			if num > res && num < k {
				res = num
			}
			left++
		} else {
			right--
		}
	}
	return res
}
