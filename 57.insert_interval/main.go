package main

import (
	"cmp"
	"slices"
)

// Brute Force like problem 56
// time complexity: O(N log N)
// space complexity: O(log N)
func insert(intervals [][]int, newInterval []int) [][]int {
	mergeIntervals := intervals
	mergeIntervals = append(mergeIntervals, newInterval)

	slices.SortFunc(mergeIntervals, func(i, j []int) int {
		return cmp.Compare(i[0], j[0])
	})

	result := [][]int{mergeIntervals[0]}
	for i := 1; i < len(mergeIntervals); i++ {
		curr := mergeIntervals[i]
		lastMerge := result[len(result)-1]

		if curr[0] <= lastMerge[1] {
			lastMerge[1] = max(lastMerge[1], curr[1])
		} else {
			result = append(result, curr)
		}
	}

	return result
}

// time complexity: O(N)
// space complexity: O(1). we only use the result array to store output, so this could be considered O(1).
func insert2(intervals [][]int, newInterval []int) [][]int {
	var i int
	var result [][]int
	n := len(intervals)

	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	result = append(result, newInterval)

	for ; i < n; i++ {
		result = append(result, intervals[i])
	}

	return result
}
