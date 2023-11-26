package main

import "golang.org/x/exp/slices"

// brute force
func countPairs(nums []int, target int) int {
	var res int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] < target {
				res++
			}
		}
	}
	return res
}

func countPairs2(nums []int, target int) int {
	slices.Sort(nums)
	var res int
	for left, right := 0, len(nums)-1; left < right; {
		if nums[left]+nums[right] < target {
			res += right - left
			left++
		} else {
			right--
		}
	}
	return res
}
