package main

import "slices"

// time complexity: O(N × 2ᴺ)
// space complexity: O(N)
func subsets(nums []int) [][]int {
	result := [][]int{{}}

	var dfs func([]int, int)
	dfs = func(comb []int, start int) {
		if start >= len(nums) {
			return
		}

		for i := start; i < len(nums); i++ {
			newComb := append(slices.Clone(comb), nums[i])
			result = append(result, newComb)
			dfs(newComb, i+1)
		}
	}

	dfs([]int{}, 0)

	return result
}
