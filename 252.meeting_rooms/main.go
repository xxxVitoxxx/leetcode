package main

import (
	"cmp"
	"slices"
)

// time complexity: O(N log N)
// space complexity: O(log N)
func canAttendMeetings(intervals [][]int) bool {
	slices.SortFunc(intervals, func(i, j []int) int {
		return cmp.Compare(i[0], j[0])
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}

	return true
}
