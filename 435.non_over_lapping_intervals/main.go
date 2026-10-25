package main

import (
	"cmp"
	"slices"
)

// Greedy
// time complexity: O(N log N)
// space complexity: O(log N)
func eraseOverlapIntervals(intervals [][]int) int {
	slices.SortFunc(intervals, func(i, j []int) int {
		return cmp.Compare(i[1], j[1])
	})

	var count int
	prevEnd := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < prevEnd {
			count++
		} else {
			prevEnd = intervals[i][1]
		}
	}
	return count
}
