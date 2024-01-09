package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// time complexity: O(n)
// space complexity: O(1)
func oddEvenList(head *ListNode) *ListNode {
	var oddHead, oddTail, evenHead, evenTail *ListNode
	var isEven bool
	for head != nil {
		next := head.Next
		head.Next = nil
		if !isEven {
			if oddHead == nil {
				oddHead = head
				oddTail = head
			} else {
				oddTail.Next = head
				oddTail = head
			}
		} else {
			if evenHead == nil {
				evenHead = head
				evenTail = head
			} else {
				evenTail.Next = head
				evenTail = head
			}
		}

		isEven = !isEven
		head = next
	}
	oddTail.Next = evenHead
	return oddHead
}

// time complexity: O(log n)
// space complexity: O(1)
func oddEvenList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummyEven := &ListNode{Next: head.Next}
	odd, even := head, head.Next
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = dummyEven.Next
	return head
}
