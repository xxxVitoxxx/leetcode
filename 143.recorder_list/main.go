package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// splitting the head into two ListNode
	current := slow.Next
	slow.Next = nil

	// reverse current
	var prev *ListNode
	for current != nil {
		current.Next, prev, current = prev, current, current.Next
	}

	// merge two ListNode
	list1, list2 := head, prev
	for list1 != nil && list2 != nil {
		tmp1, tmp2 := list1.Next, list2.Next
		list1.Next = list2
		list2.Next = tmp1
		list1, list2 = tmp1, tmp2
	}
}
