package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs and recursion
// time complexity: O(n)
// space complexity: O(log n) = O(h) to keep the recursion stack, where h is a tree height.
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

// bfs and queue
// time complexity: O(n) has to visit each node.
// space complexity: O(n) the number of the leaf node or the last full nodes.
func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue, ans := []*TreeNode{root}, 0
	for len(queue) > 0 {
		element := queue[0]
		queue = queue[1:]
		ans++
		if element.Left != nil {
			queue = append(queue, element.Left)
		}
		if element.Right != nil {
			queue = append(queue, element.Right)
		}
	}
	return ans
}

// time complexity: O(h)
// space complexity: O(h)
// 當樹是完全二元樹時，高度 h 和節點數 n 之間的關係是 n = 2^h -1，時間和空間複雜度可以視為 O(log n)。
func countNodes3(root *TreeNode) int {
	h := getHeight(root)
	if h == 0 {
		return 0
	}

	var ans int
	for h > 0 {
		if getHeight(root.Right) == h-1 {
			// the height of the left subtree is equal with the height of the right subtree
			ans += 1 << (h - 1)
			root = root.Right
		} else {
			// the difference between the height of the left subtree and the height of the right subtree is 1
			ans += 1 << (h - 2)
			root = root.Left
		}
		h--
	}
	return ans
}

// the same login with countNodes3
// func countNodes(root *TreeNode) int {
//     if root == nil {
//         return 0
//     }

//     var ans int
//     h := getHeight(root.Left)
//     for h >= 0 {
//         left := getHeight(root.Left)
//         right := getHeight(root.Right)
//         if left == right {
//             ans += 1<<left
//             root = root.Right
//         } else {
//             ans += 1<<right
//             root = root.Left
//         }
//         h--
//     }
//     return ans
// }

// recursion
// time complexity: O(h)
// space complexity: O(h)
func countNodes4(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := getHeight(root.Left)
	right := getHeight(root.Right)
	if left == right {
		return 1<<left + countNodes(root.Right)
	}
	return 1<<right + countNodes(root.Left)
}

// time complexity: O(h) where h is height of a tree
func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + getHeight(node.Left)
}
