package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// recursion
// time complexity: O(N)
// space complexity: O(N)
func copyRandomList(head *Node) *Node {
	visited := make(map[*Node]*Node)

	var clone func(*Node) *Node
	clone = func(n *Node) *Node {
		if n == nil {
			return nil
		}

		if node, ok := visited[n]; ok {
			return node
		}

		newNode := &Node{Val: n.Val}

		visited[n] = newNode

		newNode.Next = clone(n.Next)
		newNode.Random = clone(n.Random)

		return newNode
	}

	return clone(head)
}

// iteration
// time complexity: O(N)
// space complexity: O(N)
func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}

	visited := make(map[*Node]*Node)
	getCloneNode := func(n *Node) *Node {
		if n == nil {
			return nil
		}

		if node, ok := visited[n]; ok {
			return node
		}

		newNode := &Node{Val: n.Val}
		visited[n] = newNode

		return newNode
	}

	original := head
	tail := &Node{Val: head.Val}
	visited[head] = tail

	for original != nil {
		tail.Next = getCloneNode(original.Next)
		tail.Random = getCloneNode(original.Random)

		original = original.Next
		tail = tail.Next
	}

	return visited[head]
}
