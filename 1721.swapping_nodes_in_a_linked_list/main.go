package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// time complexity: O(N)
// space complexity: O(N)
func swapNodes(head *ListNode, k int) *ListNode {
	var n int
	var firstNode, secondNode *ListNode
	for curr := head; curr != nil; curr = curr.Next {
		n++

		if secondNode != nil {
			secondNode = secondNode.Next
		}

		if n == k {
			firstNode = curr
			secondNode = head
		}

	}

	firstNode.Val, secondNode.Val = secondNode.Val, firstNode.Val
	return head
}

// follow up
// time complexity: O(N)
// space complexity: O(N)
func swapNodes2(head *ListNode, k int) *ListNode {
	var n int
	for curr := head; curr != nil; curr = curr.Next {
		n++
	}

	left, right := k, n-k+1
	if left == right {
		return head
	}

	if left > right {
		left, right = right, left
	}

	dummy := &ListNode{Next: head}

	leftPrev := dummy
	for i := 1; i < left; i++ {
		leftPrev = leftPrev.Next
	}

	rightPrev := dummy
	for i := 1; i < right; i++ {
		rightPrev = rightPrev.Next
	}

	leftNode, rightNode := leftPrev.Next, rightPrev.Next
	if leftNode == rightNode {
		return head
	}

	if leftNode.Next == rightNode {
		leftPrev.Next = rightNode
		leftNode.Next = rightNode.Next
		rightNode.Next = leftNode
	} else {
		leftPrev.Next, rightPrev.Next = rightNode, leftNode
		leftNode.Next, rightNode.Next = rightNode.Next, leftNode.Next
	}

	return dummy.Next
}
