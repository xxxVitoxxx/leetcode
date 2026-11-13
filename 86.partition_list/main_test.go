package main

import "testing"

func TestPartition(t *testing.T) {
	tests := []struct {
		name     string
		head     *ListNode
		x        int
		expected *ListNode
	}{{
		name: "example1",
		head: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 2,
							},
						},
					},
				},
			},
		},
		x: 3,
		expected: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
	}, {
		name: "example2",
		head: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
			},
		},
		x: 2,
		expected: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partition(tt.head, tt.x)
			for want := tt.expected; want != nil; want, got = want.Next, got.Next {
				if got == nil || got.Val != want.Val {
					t.Fatalf("got %v, want %v", got, want.Val)
				}
			}
		})
	}
}
