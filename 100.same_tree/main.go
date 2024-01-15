package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursion
// time complexity: O(n) since one has to visit each node.
// space complexity: O(n) in the worst case of a skewed tree, to keep a recursion stack.
func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	// the same
	// if p == nil || q == nil {
	// 	return p == q
	// }
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// bfs
// time complexity: O(n) since one has to visit each node.
// space complexity: O(n) in the worst case, where the tree is perfect fully balanced binary tree,
// since BFS will have to store at least an entire level of the tree in the queue and the last level has O(n) nodes.
func isSameTree2(p, q *TreeNode) bool {
	check := func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		return true
	}

	queueP := []*TreeNode{p}
	queueQ := []*TreeNode{q}
	for len(queueP) > 0 && len(queueQ) > 0 {
		elementP := queueP[0]
		queueP = queueP[1:]
		elementQ := queueQ[0]
		queueQ = queueQ[1:]
		if !check(elementP, elementQ) {
			return false
		}
		if elementP != nil && elementQ != nil {
			queueP = append(queueP, elementP.Left)
			queueP = append(queueP, elementP.Right)
			queueQ = append(queueQ, elementQ.Left)
			queueQ = append(queueQ, elementQ.Right)
		}
	}
	return true
}
