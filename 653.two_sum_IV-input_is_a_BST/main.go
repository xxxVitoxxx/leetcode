package main

import (
	"sort"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}

	list := &List{}
	list.getValue(root)
	sort.Ints(list.List)

	left, right := 0, len(list.List)-1
	for left < right {
		num := list.List[left] + list.List[right]
		if num > k {
			right--
		} else if num < k {
			left++
		} else {
			return true
		}
	}
	return false
}

type List struct {
	List []int
}

func (l *List) getValue(root *TreeNode) {
	if root != nil {
		l.List = append(l.List, root.Val)
	}

	if root.Left != nil {
		l.getValue(root.Left)
	}

	if root.Right != nil {
		l.getValue(root.Right)
	}
}

// use inorder and two pointers
func findTarget2(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}

	var list []int
	inorder(root, &list)

	left, right := 0, len(list)-1
	for left < right {
		num := list[left] + list[right]
		if num > k {
			right--
		} else if num < k {
			left++
		} else {
			return true
		}
	}
	return false
}

func inorder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, list)
	*list = append(*list, root.Val)
	inorder(root.Right, list)
}

// use dfs and hash table
func findTarget3(root *TreeNode, k int) bool {
	return dfs(root, make(map[int]struct{}), k)
}

func dfs(root *TreeNode, m map[int]struct{}, k int) bool {
	if root == nil {
		return false
	}

	if _, ok := m[k-root.Val]; ok {
		return true
	}
	m[root.Val] = struct{}{}
	return dfs(root.Left, m, k) || dfs(root.Right, m, k)
}

func findTarget4(root *TreeNode, k int) bool {
	list := []*TreeNode{root}
	m := make(map[int]struct{})
	for i, j := 0, 0; i < len(list); i = j {
		l := len(list)
		for ; j < l; j++ {
			if _, ok := m[k-list[j].Val]; ok {
				return true
			}

			m[list[j].Val] = struct{}{}

			if list[j].Left != nil {
				list = append(list, list[j].Left)
			}

			if list[j].Right != nil {
				list = append(list, list[j].Right)
			}
		}
	}
	return false
}
