package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// two pass
// time complexity: O(N)
// space complexity: O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var count int
	curr := head
	for curr != nil {
		count++
		curr = curr.Next
	}

	c := count - n
	if c == 0 {
		return head.Next
	}

	curr = head
	for c > 0 {
		c--

		if c == 0 {
			curr.Next = curr.Next.Next
		}

		curr = curr.Next
	}

	return head
}

// two pass
// time complexity: O(N)
// space complexity: O(1)
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	var length int
	curr := head
	for curr != nil {
		length++
		curr = curr.Next
	}

	length -= n

	dummy := &ListNode{0, head}
	curr = dummy
	for length > 0 {
		length--
		curr = curr.Next
	}

	curr.Next = curr.Next.Next

	return dummy.Next
}

// one pass
// time complexity: O(N)
// space complexity: O(1)
func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, slow := dummy, dummy
	for i := 0; i <= n; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return dummy.Next
}
