package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maximumAverageSubtree(root *TreeNode) float64 {
	var ans float64
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		cal := getSubtreeValue(node) / getSubtreeCount(node)
		if cal > ans {
			ans = cal
		}
	}
	return ans
}

func getSubtreeCount(node *TreeNode) float64 {
	if node == nil {
		return 0
	}
	return 1 + getSubtreeCount(node.Left) + getSubtreeCount(node.Right)
}

func getSubtreeValue(node *TreeNode) float64 {
	if node == nil {
		return 0
	}
	return float64(node.Val) + getSubtreeValue(node.Left) + getSubtreeValue(node.Right)
}

// use dfs
func maximumAverageSubtree2(root *TreeNode) float64 {
	var ans float64
	var dfs func(*TreeNode) (float64, float64)
	dfs = func(node *TreeNode) (float64, float64) {
		if node == nil {
			return 0, 0
		}

		leftSum, leftCount := dfs(node.Left)
		rightSum, rightCount := dfs(node.Right)
		sum := float64(node.Val) + leftSum + rightSum
		count := 1 + leftCount + rightCount

		cal := sum / count
		if cal > ans {
			ans = cal
		}
		return sum, count
	}
	dfs(root)
	return ans
}
