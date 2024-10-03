package main

import (
	"sort"
)

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, curr := range nums {
		if val, ok := m[target-curr]; ok {
			return []int{val, i}
		}
		m[curr] = i
	}
	return nil
}

// two point
// time complexity: O(n) in the worst case, where we might need to traverse the entire array once
// space complexity: O(1) for two pointers and other variables
func twoSum2(nums []int, target int) []int {
	numbersWithIndices := make([][2]int, len(nums))
	for i, num := range nums {
		numbersWithIndices[i] = [2]int{i, num}
	}

	sort.Slice(numbersWithIndices, func(i, j int) bool {
		return numbersWithIndices[i][1] < numbersWithIndices[j][1]
	})

	left, right := 0, len(nums)-1

	for left < right {
		sum := numbersWithIndices[left][1] + numbersWithIndices[right][1]
		if sum == target {
			return []int{numbersWithIndices[left][0], numbersWithIndices[right][0]}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{}
}
