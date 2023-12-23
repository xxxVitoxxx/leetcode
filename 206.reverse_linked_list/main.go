package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var rev *ListNode
	for head != nil {
		// head.Next, rev, head = rev, head, head.Next
		node := head
		head = head.Next
		node.Next = rev
		rev = node
	}
	return rev
}

func reverseList2(head *ListNode) *ListNode {
	return helper2(head, nil)
}

func helper2(current, reverse *ListNode) *ListNode {
	if current == nil {
		return reverse
	}

	next := current.Next
	current.Next = reverse
	return helper2(next, current)

}

// use recursion
func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList3(head.Next)
	nextHead := head.Next
	nextHead.Next = head
	head.Next = nil
	return newHead
}
