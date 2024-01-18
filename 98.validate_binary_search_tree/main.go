package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursive inorder traversal
// time complexity: O(n)
// space complexity: O(n)
func isValidBST(root *TreeNode) bool {
	var nums []int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Left)
			nums = append(nums, node.Val)
			dfs(node.Right)
		}
	}
	dfs(root)

	if len(nums) == 0 {
		return true
	}
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] <= prev {
			return false
		}
		prev = nums[i]
	}

	return true
}

// recursive traversal with valid range
// time complexity: O(n)
// space complexity: O(n)
func isValidBST2(root *TreeNode) bool {
	var validate func(*TreeNode, int, int) bool
	validate = func(node *TreeNode, low, high int) bool {
		if node == nil {
			return true
		}

		if low >= node.Val || high <= node.Val {
			return false
		}

		return validate(node.Left, low, node.Val) && validate(node.Right, node.Val, high)
	}

	return validate(root, -(1<<31 + 1), 1<<31)
}

// iterative inorder traversal
// time complexity: O(n)
// space complexity: O(n)
func isValidBST3(root *TreeNode) bool {
	var stack []*TreeNode
	prev := -(1<<31 + 1)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= prev {
			return false
		}

		prev = root.Val
		root = root.Right
	}
	return true
}
