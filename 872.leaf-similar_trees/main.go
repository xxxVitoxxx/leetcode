package main

import "reflect"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs
// time complexity: O(n)
// space complexity: O(l1+l2) where l1 ans l2 are leaf nodes in tree1 and tree2
func leafSimilar(root1, root2 *TreeNode) bool {
	list1, list2 := &[]int{}, &[]int{}
	dfs(root1, list1)
	dfs(root2, list2)
	return reflect.DeepEqual(*list1, *list2)
}

func dfs(root *TreeNode, list *[]int) {
	if root.Left != nil {
		dfs(root.Left, list)
	}
	if root.Right != nil {
		dfs(root.Right, list)
	}
	if root.Left == nil && root.Right == nil {
		*list = append(*list, root.Val)
	}
}

// dfs
// time complexity: O(n)
// space complexity: O(l1+l2)
func leafSimilar2(root1, root2 *TreeNode) bool {
	var dfs func(root *TreeNode) []int
	dfs = func(root *TreeNode) []int {
		if root == nil {
			return []int{}
		}
		if root.Left == nil && root.Right == nil {
			return []int{root.Val}
		}
		return append(dfs(root.Left), dfs(root.Right)...)
	}
	return reflect.DeepEqual(dfs(root1), dfs(root2))
}
