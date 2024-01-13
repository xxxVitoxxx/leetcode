package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// bfs and queue
// time complexity: O(n) since one has to visit each node.
// space complexity: O(d) where d is a tree diameter.
func rightSideView(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		level := len(queue)
		ans = append(ans, queue[level-1].Val)
		for i := 0; i < level; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[level:]
	}
	return ans
}

// dfs
// time complexity: O(n) since one has to visit each node.
// space complexity: O(h) to keep the recursion stack, where h is a tree height.
// the worst case situation is a skewed tree when h = n.
func rightSideView2(root *TreeNode) []int {
	var ans []int
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level == len(ans) {
			ans = append(ans, node.Val)
		}
		if node.Right != nil {
			dfs(node.Right, level+1)
		}
		if node.Left != nil {
			dfs(node.Left, level+1)
		}
	}
	dfs(root, 0)
	return ans
}
