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
func maxLevelSum(root *TreeNode) int {
	maximumLevel, maximumSum, level := 1, root.Val, 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		level++
		size, levelSum := len(queue), 0
		for i := 0; i < size; i++ {
			levelSum += queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		queue = queue[size:]
		if levelSum > maximumSum {
			maximumSum = levelSum
			maximumLevel = level
		}
	}
	return maximumLevel
}

// dfs
// time complexity: O(n)
// space complexity: O(n)
func maxLevelSum2(root *TreeNode) int {
	var levelSum []int
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if len(levelSum) == level {
			levelSum = append(levelSum, node.Val)
		} else {
			levelSum[level] += node.Val
		}

		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 0)

	maximumLevel, maximumSum := 0, levelSum[0]
	for level, sum := range levelSum {
		if sum > maximumSum {
			maximumSum = sum
			maximumLevel = level
		}
	}
	return maximumLevel + 1
}
