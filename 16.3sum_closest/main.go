package main

import (
	"sort"
)

// brute force
func threeSumClosest(nums []int, target int) int {
	res, diff := 0, 1<<63-1
	for first := 0; first < len(nums)-2; first++ {
		for second := first + 1; second < len(nums)-1; second++ {
			for third := second + 1; third < len(nums); third++ {
				num := nums[first] + nums[second] + nums[third]
				if num == target {
					return num
				}

				d := abs(target - num)
				if d < diff {
					diff = d
					res = num
				}
			}
		}
	}
	return res
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

// use two pointer
func threeSumClosest2(nums []int, target int) int {
	sort.Ints(nums)

	res, diff := 0, 1<<63-1
	for first := 0; first < len(nums)-2; first++ {
		if first > 0 && nums[first-1] == nums[first] {
			continue
		}

		second, third := first+1, len(nums)-1
		for second < third {
			num := nums[first] + nums[second] + nums[third]
			d := abs(target - num)
			if diff > d {
				diff = d
				res = num
			}

			if num > target {
				third--
				for second < third && nums[third] == nums[third+1] {
					third--
				}
			} else if num < target {
				second++
				for second < third && nums[second] == nums[second-1] {
					second++
				}
			} else {
				return num
			}
		}
	}
	return res
}
