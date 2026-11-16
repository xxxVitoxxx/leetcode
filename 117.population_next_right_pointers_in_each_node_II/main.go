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
		return nil
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

// time complexity: O(N)
// space complexity: O(1)
func connect2(root *Node) *Node {
	dummy := new(Node)
	curr := root

	for curr != nil {
		dummy.Next = nil
		tail := dummy

		for n := curr; n != nil; n = n.Next {
			if n.Left != nil {
				tail.Next = n.Left
				tail = tail.Next
			}

			if n.Right != nil {
				tail.Next = n.Right
				tail = tail.Next
			}
		}

		curr = dummy.Next
	}

	return root
}
