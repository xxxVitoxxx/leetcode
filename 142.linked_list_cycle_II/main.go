package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}

func detectCycle2(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			l := cycleLen(slow)
			pnt1, pnt2 := head, head
			for l > 0 {
				pnt1 = pnt1.Next
				l--
			}

			for pnt1 != pnt2 {
				pnt1 = pnt1.Next
				pnt2 = pnt2.Next
			}
			return pnt1
		}
	}
	return nil
}

func cycleLen(node *ListNode) int {
	var l int
	curry := node
	for curry != nil {
		l++
		curry = curry.Next
		if curry == node {
			return l
		}
	}
	return 0
}
