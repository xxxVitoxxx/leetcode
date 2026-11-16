package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// time complexity: O(N)
// space complexity: O(N)
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		size := len(queue)

		for i := range size {
			curr := queue[i]
			if i+1 < size {
				curr.Next = queue[i+1]
			}

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}

		queue = queue[size:]
	}

	return root
}

// using previously established next pointers
// time complexity: O(N)
// space complexity: O(1)
func connect2(root *Node) *Node {
	if root == nil {
		return nil
	}

	leftmost := root
	for leftmost.Left != nil {
		head := leftmost
		for head != nil {
			head.Left.Next = head.Right

			if head.Next != nil {
				head.Right.Next = head.Next.Left
			}

			head = head.Next
		}

		leftmost = leftmost.Left
	}

	return root
}
