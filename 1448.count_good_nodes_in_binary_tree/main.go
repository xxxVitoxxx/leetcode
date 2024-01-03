package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs
// time complexity: O(n) where n is the number of nodes in the binary tree
// space complexity: O(h) where h is the height of the binary tree.
// in the worst case(for a skewed tree), it can be O(n)
func goodNodes(root *TreeNode) int {
	var count int
	var dfs func(root *TreeNode, max int)
	dfs = func(root *TreeNode, max int) {
		if root == nil {
			return
		}
		if root.Val >= max {
			count++
			max = root.Val
		}
		dfs(root.Left, max)
		dfs(root.Right, max)
	}

	dfs(root, root.Val)

	return count
}

// dfs
// time complexity: O(n)
// space complexity: O(h)
func goodNodes2(root *TreeNode) int {
	var dfs func(root *TreeNode, max int) int
	dfs = func(root *TreeNode, max int) int {
		if root == nil {
			return 0
		}
		if root.Val >= max {
			max = root.Val
			return dfs(root.Left, max) + dfs(root.Right, max) + 1
		}
		return dfs(root.Left, max) + dfs(root.Right, max)
	}
	return dfs(root, root.Val)
}
