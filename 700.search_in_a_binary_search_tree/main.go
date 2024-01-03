package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// use recursion
// time complexity: O(log n) is tree height. in the worst case(for a skewed tree), it can be O(n).
// space complexity: O(log n) to keep recursion stack. in the worst case(for a skewed tree), it can be O(n).
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)
}

// use iteration
// time complexity: O(n)
// space complexity: O(1) since it's a constant space solution.
func searchBST2(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if root.Val < val {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return root
}
