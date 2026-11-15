package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// time complexity: O(N)
// space complexity: O(1)
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	origin := dummy
	for i := 1; i < left; i++ {
		origin = origin.Next
	}

	tail := origin.Next
	curr := tail
	var reversed *ListNode
	for i := left; i <= right; i++ {
		next := curr.Next
		curr.Next = reversed
		reversed = curr
		curr = next
	}

	tail.Next = curr
	origin.Next = reversed

	return dummy.Next
}
