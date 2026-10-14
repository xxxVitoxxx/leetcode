package main

import "sort"

// time complexity: O(N log N)
// space complexity: O(log N)
func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1]
	})

	var result int

	n := len(costs) / 2
	for i := range n {
		result += costs[i][0] + costs[i+n][1]
	}

	return result
}
