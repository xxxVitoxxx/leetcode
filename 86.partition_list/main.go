package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// time complexity: O(N)
// space complexity: O(1). we have not utilized extra space.
func partition(head *ListNode, x int) *ListNode {
	lessNodes := new(ListNode)
	less := lessNodes
	geNodes := new(ListNode)
	ge := geNodes

	for curr := head; curr != nil; curr = curr.Next {
		if curr.Val < x {
			less.Next = curr
			less = less.Next
		} else {
			ge.Next = curr
			ge = ge.Next
		}
	}

	ge.Next = nil
	less.Next = geNodes.Next
	return lessNodes.Next
}
