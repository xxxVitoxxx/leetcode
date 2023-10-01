package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// stack
func isPalindrome(head *ListNode) bool {
	var stack []int
	for node := head; node != nil; node = node.Next {
		stack = append(stack, node.Val)
	}

	for node := head; node != nil; node = node.Next {
		if node.Val != stack[len(stack)-1] {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return true
}

// fast-slow pointers
func isPalindrome2(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	var prev *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	for head != nil && prev != nil {
		if head.Val != prev.Val {
			return false
		}
		head = head.Next
		prev = prev.Next
	}
	return true
}
