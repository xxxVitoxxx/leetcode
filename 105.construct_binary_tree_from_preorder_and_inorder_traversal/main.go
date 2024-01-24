package main

import "golang.org/x/exp/slices"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// time complexity: O(n^2) where n is the number of nodes in the binary tree.
// in each recursive call, the algorithm search for the root value index in the inorder slice,
// resulting in a total time complexity of O(n^2)
// space complexity: O(n)
func buildTree(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	i := slices.Index(inorder, preorder[0])
	lpre, rpre := preorder[1:i+1], preorder[i+1:]
	lin, rin := inorder[:i], inorder[i+1:]
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(lpre, lin),
		Right: buildTree(rpre, rin),
	}
}

// optimized to O(n) by using hashmap to store indices
// time complexity: O(n)
// space complexity: O(n)
func buildTree2(preorder, inorder []int) *TreeNode {
	inorderMap := make(map[int]int, len(inorder))
	for i := 0; i < len(inorder); i++ {
		inorderMap[inorder[i]] = i
	}

	var preorderIndex int
	var arrToTree func(int, int) *TreeNode
	arrToTree = func(left, right int) *TreeNode {
		if left > right || preorderIndex == len(preorder) {
			return nil
		}

		root := &TreeNode{}
		root.Val = preorder[preorderIndex]
		preorderIndex++

		root.Left = arrToTree(left, inorderMap[root.Val]-1)
		root.Right = arrToTree(inorderMap[root.Val]+1, right)

		return root
	}

	return arrToTree(0, len(preorder)-1)
}
