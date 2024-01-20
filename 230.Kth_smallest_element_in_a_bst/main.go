package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursive inorder traversal
// time complexity: O(n)
// sapce complexity: O(n)
func kthSmallest(root *TreeNode, k int) int {
	var ans int
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node != nil {
			inorder(node.Left)
			k--
			if k == 0 {
				ans = node.Val
				return
			}
			inorder(node.Right)
		}
	}
	inorder(root)
	return ans
}

// iterative inorder traversal
// time complexity: O(h) = O(log(n)) where h is a tree height. this complexity is defined by the stack.
// space complexity: O(h) = O(log(n)) to keep the stack.
func kthSmallest2(root *TreeNode, k int) int {
	var ans int
	var stack []*TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			ans = root.Val
			break
		}
		root = root.Right
	}
	return ans
}
