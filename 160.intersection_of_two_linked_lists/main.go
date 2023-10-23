package main

// Definition of singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// use hash table
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	for node := headA; node != nil; node = node.Next {
		m[node] = struct{}{}
	}
	for node := headB; node != nil; node = node.Next {
		if _, ok := m[node]; ok {
			return node
		}
	}
	return nil
}

// use two pointers
// if two lists intersect, their intersection point can be found
// by changing headA to headB at the end of the path.
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a != nil {
			a = a.Next
		} else {
			a = headB
		}

		if b != nil {
			b = b.Next
		} else {
			b = headA
		}
	}
	return a
}

func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	var lenA int
	var lenB int
	for node := headA; node != nil; node = node.Next {
		lenA++
	}
	for node := headB; node != nil; node = node.Next {
		lenB++
	}

	currA, currB := headA, headB
	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			currA = currA.Next
		}
	}
	if lenA < lenB {
		for i := 0; i < lenB-lenA; i++ {
			currB = currB.Next
		}
	}

	for currA != nil && currB != nil {
		if currA == currB {
			return currA
		}
		currA, currB = currA.Next, currB.Next
	}
	return nil
}
