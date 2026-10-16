package main

import "slices"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Brute Force
// time complexity: O(N log N)
// space complexity: O(N)
func mergeKLists(lists []*ListNode) *ListNode {
	var values []int
	for _, list := range lists {
		for list != nil {
			values = append(values, list.Val)
			list = list.Next
		}
	}

	slices.Sort(values)

	dummy := &ListNode{}
	curr := dummy
	for _, value := range values {
		curr.Next = &ListNode{Val: value}
		curr = curr.Next
	}
	return dummy.Next
}

// N is the number of linked lists, and k is the average length of each linked list.
// time complexity: O(N * K)
// space complexity: O(1)
func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	merge := lists[0]

	for i := 1; i < len(lists); i++ {
		merge = mergeTwoList(merge, lists[i])
	}

	return merge
}

func mergeTwoList(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}

		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	} else if list2 != nil {
		curr.Next = list2
	}

	return dummy.Next
}
