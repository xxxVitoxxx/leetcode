package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// time complexity: O(N)
// space complexity: O(N). stack space utilized for recursion.
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}

// time complexity: O(N)
// space complexity: O(1)
func swapPairs2(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for curr := head; curr != nil && curr.Next != nil; {
		firstNode := curr
		secondNode := curr.Next

		prev.Next = secondNode
		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode

		prev = firstNode
		curr = firstNode.Next
	}

	return dummy.Next
}
