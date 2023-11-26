package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}

	l, tail, sep := 1, head, head
	for ; tail.Next != nil; tail = tail.Next {
		l++
	}

	if k %= l; k == 0 {
		return head
	}
	for i := 1; i < l-k; i++ {
		sep = sep.Next
	}

	tail.Next = head
	newHead := sep.Next
	sep.Next = nil

	return newHead
}
