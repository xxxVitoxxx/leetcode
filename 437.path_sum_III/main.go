package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// stack and dfs
func pathSum(root *TreeNode, targetSum int) int {
	var ans int
	var dfs func(root *TreeNode, path []int)
	dfs = func(root *TreeNode, path []int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		var sum int
		for i := len(path) - 1; i >= 0; i-- {
			sum += path[i]
			if sum == targetSum {
				ans++
			}
		}
		dfs(root.Left, path)
		dfs(root.Right, path)
	}
	dfs(root, []int{})
	return ans
}
