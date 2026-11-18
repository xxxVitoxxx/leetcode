package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// time complexity: O(N)
// space complexity: O(N)
func diameterOfBinaryTree(root *TreeNode) int {
	var result int
	var dfs func(*TreeNode) int
	dfs = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		left := dfs(n.Left)
		right := dfs(n.Right)

		result = max(result, left+right)

		return max(left, right) + 1
	}

	dfs(root)

	return result
}
