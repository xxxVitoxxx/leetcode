package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// iterative approach
// time complexity: O(n + m). because exactly one of list1 and list2 is incremented on each loop iteration,
// the while loop runs for a number of iterations equal to the sum of the lengths of the two lists.
// all other work is constant, so the overall complexity is linear.
//
// space complexity: O(1)
func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	result := &ListNode{}
	dummy := result
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			dummy.Next = list1
			list1 = list1.Next
		} else {
			dummy.Next = list2
			list2 = list2.Next
		}
		dummy = dummy.Next
	}

	if list1 != nil {
		dummy.Next = list1
	} else {
		dummy.Next = list2
	}

	return result.Next
}

// recursive approach
// time complexity: O(n + m)
// space complexity: O(n + m). the first call to mergeTwoLists2 does not return until the ends
// of both list1 and list2 have been reached, so n + m stack frames consume O(n + m) space.
func mergeTwoLists2(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = mergeTwoLists2(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists2(list1, list2.Next)
		return list2
	}
}
