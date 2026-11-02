package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS
// time complexity: O(N)
// space complexity: O(N)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

// BFS
// time complexity: O(N)
// time complexity: O(1)
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	node := root
	for node != nil {
		if node.Val > p.Val && node.Val > q.Val {
			node = node.Left
		} else if node.Val < p.Val && node.Val < q.Val {
			node = node.Right
		} else {
			break
		}
	}
	return node
}
