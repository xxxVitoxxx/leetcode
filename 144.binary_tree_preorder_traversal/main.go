package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// time complexity: O(N)
// space complexity: O(N)
func preorderTraversal(root *TreeNode) []int {
	var result []int
	var traversal func(*TreeNode)
	traversal = func(n *TreeNode) {
		if n == nil {
			return
		}

		result = append(result, n.Val)
		traversal(n.Left)
		traversal(n.Right)
	}

	traversal(root)

	return result
}
