package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left) + 1
	right := maxDepth(root.Right) + 1
	var max int
	if left > right {
		max = left
	} else {
		max = right
	}
	return max
}
