package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs
// time complexity: O(n)
// space complexity: O(d) where d is depth of the tree
func longestZigZag(root *TreeNode) int {
	var dfs func(node *TreeNode, isLeft bool, count int) int
	dfs = func(node *TreeNode, isLeft bool, count int) int {
		if node == nil {
			return count
		}

		if isLeft {
			return max(dfs(node.Left, true, 0), dfs(node.Right, false, count+1))
		}
		return max(dfs(node.Left, true, count+1), dfs(node.Right, false, 0))
	}

	return max(dfs(root.Left, true, 0), dfs(root.Right, false, 0))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dfs
// time complexity: O(n)
// space complexity: O(d) where d is depth of the tree
func longestZigZag2(root *TreeNode) int {
	var ans int
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, left, right int) {
		if node != nil {
			ans = max2(ans, left, right)
			dfs(node.Left, right+1, 0)
			dfs(node.Right, 0, left+1)
		}
	}
	dfs(root, 0, 0)
	return ans
}

func max2(a ...int) int {
	nums, max := a, 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
