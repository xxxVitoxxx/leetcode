package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// two pointer
// time complexity: O(n)
// space complexity: O(1)
func pairSum(head *ListNode) int {
	fast, slow := head, head
	// get middle of the linked list
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var reverse *ListNode
	for slow != nil {
		slow, slow.Next, reverse = slow.Next, reverse, slow
	}

	var ans int
	for reverse != nil {
		if head.Val+reverse.Val > ans {
			ans = head.Val + reverse.Val
		}
		head = head.Next
		reverse = reverse.Next
	}
	return ans
}

// arr or stack
// time complexity: O(n)
// space complexity: O(n)
func pairSum2(head *ListNode) int {
	var arr []int
	for node := head; node != nil; node = node.Next {
		arr = append(arr, node.Val)
	}

	var ans int
	left, right := 0, len(arr)-1
	for left < right {
		if arr[left]+arr[right] > ans {
			ans = arr[left] + arr[right]
		}
		left, right = left+1, right-1
	}
	return ans
}
