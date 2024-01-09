package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursion
// time complexity: O(h) where h is the height of the binary search tree.
// space complexity: O(h) where h is the height of the binary search tree.
// in the worst case(for a skewed tree), it can be O(n)
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			// if the node has both left and right children, find the minium node in the right subtree.
			// replace the current node's value with that minium value.
			root.Val = findMin(root.Right)
			// recursively delete minium node in the right subtree.
			root.Right = deleteNode(root.Right, root.Val)
		}
	}
	return root
}

func findMin(node *TreeNode) int {
	minVal := node.Val
	for node != nil {
		if node.Val < minVal {
			minVal = node.Val
		}
		node = node.Left
	}
	return minVal
}

// time complexity: O(h)
// time complexity: O(h)
func deleteNode2(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			// if the node has both left and right children
			// find the minium node in the right subtree and replace it.
			minSubTree := findMinNode(root.Right)
			leftSubTree := root.Left
			rightSubTree := deleteNode(root.Right, minSubTree.Val)
			minSubTree.Left = leftSubTree
			minSubTree.Right = rightSubTree
			return minSubTree
		}
	}

	return root
}

func findMinNode(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Left == nil {
		return node
	}
	return findMinNode(node.Left)
}
