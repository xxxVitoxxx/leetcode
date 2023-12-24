package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	var l int
	for node := head; node != nil; node = node.Next {
		l++
	}
	if l < 2 {
		return nil
	}

	i, middle := 0, l/2
	var prev *ListNode
	for node := head; node != nil; node = node.Next {
		if i == middle {
			prev.Next = node.Next
			break
		}
		i++
		prev = node
	}
	return head
}

func deleteMiddle2(head *ListNode) *ListNode {
	var l int
	for node := head; node != nil; node = node.Next {
		l++
	}
	if l < 2 {
		return nil
	}

	node := head
	for i := 0; i < l/2-1; i++ {
		node = node.Next
	}
	node.Next = node.Next.Next
	return head
}

// use slow fast pointer
func deleteMiddle3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	var slowPrev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slowPrev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	slowPrev.Next = slow.Next
	return head
}

// use slow fast pointer
func deleteMiddle4(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	slow, fast := dummy, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
