package main

import "sort"

func sortedSquares(nums []int) []int {
	for i, num := range nums {
		nums[i] = pow(abs(num))
	}
	sort.Ints(nums)
	return nums
}

func sortedSquares2(nums []int) []int {
	res := make([]int, len(nums))
	left, right := 0, len(nums)-1
	for i := len(res) - 1; i >= 0; i-- {
		leftNum, rightNum := abs(nums[left]), abs(nums[right])
		if leftNum > rightNum {
			res[i] = pow(leftNum)
			left++
		} else {
			res[i] = pow(rightNum)
			right--
		}
	}
	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func pow(n int) int {
	return n * n
}

func sortedSquares3(nums []int) []int {
	res := make([]int, len(nums))
	left, right := 0, len(nums)-1
	for i := len(res) - 1; i >= 0; i-- {
		if nums[left] < 0 && -nums[left] > nums[right] {
			res[i] = -nums[left] * -nums[left]
			left++
		} else {
			res[i] = nums[right] * nums[right]
			right--
		}
	}
	return res
}

func sortedSquares4(nums []int) []int {
	for i, num := range nums {
		nums[i] = num * num
	}
	sort.Ints(nums)
	return nums
}
