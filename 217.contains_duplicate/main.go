package main

import "sort"

// use hash
func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = struct{}{}
	}
	return false
}

// brute force
func containsDuplicate2(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func containsDuplicate3(nums []int) bool {
	sort.Ints(nums)
	for i := range nums {
		if i > 0 && nums[i-1] == nums[i] {
			return true
		}
	}
	return false
}
