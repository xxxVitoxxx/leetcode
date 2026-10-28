package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS
// time complexity: O(MN)
// for every N node in the tree, we check if the tree rooted at node is identical to subRoot.
// this check takes O(M) time, where M is the number of nodes in subRoot
//
// space complexity: O(M+N)
// there will be at most N recursive call to dfs. each of these calls will have M recursive calls to isIdentical.
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	return dfs(root, subRoot)
}

func dfs(node, subRoot *TreeNode) bool {
	if node == nil {
		return false
	}

	if isIdentical(node, subRoot) {
		return true
	}

	return dfs(node.Left, subRoot) || dfs(node.Right, subRoot)
}

func isIdentical(node1, node2 *TreeNode) bool {
	if node1 == nil || node2 == nil {
		return node1 == nil && node2 == nil
	}

	return node1.Val == node2.Val && isIdentical(node1.Left, node2.Left) && isIdentical(node1.Right, node2.Right)
}

// String Match
// time complexity: O(M+N)
// space complexity: O(M+N)
func isSubtree2(root *TreeNode, subRoot *TreeNode) bool {
	return strings.Contains(serialize(root), serialize(subRoot))
}

func serialize(root *TreeNode) string {
	var str strings.Builder
	var dfs func(*TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			str.WriteString(",#")
			return
		}

		str.WriteByte(',')
		str.WriteString(strconv.Itoa(n.Val))

		dfs(n.Left)
		dfs(n.Right)
	}

	dfs(root)

	return str.String()
}
