package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// time complexity: O(N)
// space complexity: O(1)
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	curr := head
	for curr != nil {
		if curr.Next != nil && curr.Val == curr.Next.Val {
			for curr.Next != nil && curr.Val == curr.Next.Val {
				curr = curr.Next
			}

			prev.Next = curr.Next
		} else {
			prev = prev.Next
		}

		curr = curr.Next
	}
	return dummy.Next
}
