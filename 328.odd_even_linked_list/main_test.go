package main

import (
	"testing"
)

func TestOddEvenList(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	list2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 7,
							},
						},
					},
				},
			},
		},
	}

	tests := []struct {
		name   string
		head   *ListNode
		ansVal []int
	}{
		{"example1", list1, []int{1, 3, 5, 2, 4}},
		{"example2", list2, []int{2, 3, 6, 7, 1, 5, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := oddEvenList(tt.head)
			for i := 0; res != nil; i++ {
				if res.Val != tt.ansVal[i] {
					t.Errorf("got %d, want %d", res.Val, tt.ansVal[i])
				}
				res = res.Next
			}
		})
	}
}

func TestOddEvenList2(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	list2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 7,
							},
						},
					},
				},
			},
		},
	}

	tests := []struct {
		name   string
		head   *ListNode
		ansVal []int
	}{
		{"example1", list1, []int{1, 3, 5, 2, 4}},
		{"example2", list2, []int{2, 3, 6, 7, 1, 5, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := oddEvenList2(tt.head)
			for i := 0; res != nil; i++ {
				if res.Val != tt.ansVal[i] {
					t.Errorf("got %d, want %d", res.Val, tt.ansVal[i])
				}
				res = res.Next
			}
		})
	}
}
