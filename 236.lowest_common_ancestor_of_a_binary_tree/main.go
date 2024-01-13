package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs
// time complexity: O(n)
// space complexity: O(n)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	var res *TreeNode
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}

		self := node == p || node == q
		left := dfs(node.Left)
		right := dfs(node.Right)
		if self && (left || right) || left && right {
			res = node
		}

		return self || left || right
	}
	dfs(root)
	return res
}
