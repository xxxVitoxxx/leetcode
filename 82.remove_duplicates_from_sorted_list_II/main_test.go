package main

import "testing"

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		head     *ListNode
		expected *ListNode
	}{{
		name: "example1",
		head: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
			},
		},
		expected: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 5,
				},
			},
		},
	}, {
		name: "example2",
		head: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
		},
		expected: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteDuplicates(tt.head)
			for want := tt.expected; want != nil; want, got = want.Next, got.Next {
				if got == nil || got.Val != want.Val {
					t.Errorf("got %v, want %v", got.Val, want.Val)
				}
			}
		})
	}
}
