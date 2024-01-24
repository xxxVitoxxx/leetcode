package main

// time complexity: O(1)
// space complexity: O(1)
func findCenter(edges [][]int) int {
	switch edges[0][0] {
	case edges[1][0]:
		return edges[0][0]
	case edges[1][1]:
		return edges[0][0]
	}
	return edges[0][1]
}
