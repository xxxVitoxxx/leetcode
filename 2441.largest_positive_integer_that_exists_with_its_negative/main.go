package main

import "sort"

// brute force
func findMaxK(nums []int) int {
	m := make(map[int]struct{})
	check := func(n int) bool {
		if _, ok := m[n]; ok {
			return true
		}
		return false
	}

	res := -1
	for _, num := range nums {
		if num < 0 {
			if check(-num) && -num > res {
				res = -num
			}
		} else {
			if check(-num) && num > res {
				res = num
			}
		}
		m[num] = struct{}{}
	}
	return res
}

func findMaxK2(nums []int) int {
	mapPos := make(map[int]struct{})
	mapNeg := make(map[int]struct{})
	for _, num := range nums {
		if num > 0 {
			mapPos[num] = struct{}{}
		} else {
			mapNeg[num] = struct{}{}
		}
	}

	res := -1
	for k := range mapPos {
		if _, ok := mapNeg[-k]; ok && k > res {
			res = k
		}
	}
	return res
}

func findMaxK3(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		m[num] = struct{}{}
	}

	sort.Ints(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		if _, ok := m[-nums[i]]; ok {
			return nums[i]
		}
	}
	return -1
}

func findMaxK4(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	res := -1
	for _, num := range nums {
		abs := abs(num)
		if _, ok := m[-num]; ok && abs > res {
			res = abs
		}
		m[num] = struct{}{}
	}
	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// use two pointer
func findMaxK5(nums []int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left < right {
		if -nums[left] == nums[right] {
			return nums[right]
		}

		if -nums[left] > nums[right] {
			left++
		} else {
			right--
		}
	}
	return -1
}
