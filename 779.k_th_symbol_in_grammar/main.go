package main

import "math"

// binary tree traversal
// time complexity: O(n)
// with each recursive call, we reduce n by one util n becomes equal to 1.
// as a result, the overall time complexity is O(n)
//
// space complexity: O(n)
// each recursive call will add a new frame to the stack util we reach the base case.
// (when n util becomes equal to 1). hence, the space complexity is O(n)
func kthGrammar(n, k int) int {
	return dfs(n, k, 0)
}

func dfs(n, k, rootValue int) int {
	if n == 1 {
		return rootValue
	}

	totalNodes := int(math.Pow(2, float64(n-1)))
	half := totalNodes / 2

	if k > half {
		var nextRootValue int
		if rootValue == 0 {
			nextRootValue = 1
		}
		return dfs(n-1, k-half, nextRootValue)
	} else {
		var nextRootValue int
		if rootValue == 1 {
			nextRootValue = 1
		}
		return dfs(n-1, k, nextRootValue)
	}
}
