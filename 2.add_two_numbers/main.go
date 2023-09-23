package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	carry, dummy := 0, &ListNode{}
	for node := dummy; l1 != nil || l2 != nil || carry > 0; node = node.Next {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		node.Next = &ListNode{Val: carry % 10}
		carry /= 10
	}
	return dummy.Next
}
