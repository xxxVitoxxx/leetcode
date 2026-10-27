package main

import (
	"cmp"
	"slices"
)

// time complexity: O(N log N)
// space complexity: O(log N)
func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(i, j []int) int {
		return cmp.Compare(i[0], j[0])
	})

	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]

		lastMerge := result[len(result)-1]
		if curr[0] <= lastMerge[1] {
			lastMerge[1] = max(curr[1], lastMerge[1])
		} else {
			result = append(result, curr)
		}
	}

	return result
}
