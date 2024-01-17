package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// bfs
// time complexity: O(n)
// space complexity: O(n)
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ans [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, size)
		for i := 0; i < size; i++ {
			level[i] = queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		queue = queue[size:]
		ans = append(ans, level)
	}
	return ans
}

// dfs
// time complexity: O(n)
// space complexity: O(n)
func levelOrder2(root *TreeNode) [][]int {
	var ans [][]int
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if len(ans) == level {
			ans = append(ans, [][]int{{node.Val}}...)
		} else {
			ans[level] = append(ans[level], node.Val)
		}

		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 0)
	return ans
}
