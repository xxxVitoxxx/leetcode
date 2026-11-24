package main

import "slices"

// time complexity: O(N × 2ᴺ)
// space complexity: O(N)
func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)

	result := [][]int{}
	path := []int{}
	var dfs func(int)
	dfs = func(start int) {
		result = append(result, slices.Clone(path))

		for i := start; i < len(nums); i++ {
			if i != start && nums[i-1] == nums[i] {
				continue
			}

			path = append(path, nums[i])
			dfs(i + 1)

			path = path[:len(path)-1]
		}
	}

	dfs(0)

	return result
}
