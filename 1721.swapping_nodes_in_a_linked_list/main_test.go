package main

import (
	"testing"
)

func TestSwapNodes(t *testing.T) {
	tests := []struct {
		name     string
		head     *ListNode
		k        int
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
						Val: 4,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		},
		k: 2,
		expected: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		},
	}, {
		name: "example2",
		head: &ListNode{
			Val: 7,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 7,
							Next: &ListNode{
								Val: 8,
								Next: &ListNode{
									Val: 3,
									Next: &ListNode{
										Val: 0,
										Next: &ListNode{
											Val: 9,
											Next: &ListNode{
												Val: 5,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		k: 5,
		expected: &ListNode{
			Val: 7,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 8,
							Next: &ListNode{
								Val: 7,
								Next: &ListNode{
									Val: 3,
									Next: &ListNode{
										Val: 0,
										Next: &ListNode{
											Val: 9,
											Next: &ListNode{
												Val: 5,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}, {
		name: "example3",
		head: &ListNode{
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
		},
		k: 5,
		expected: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := swapNodes(cloneNode(tt.head), tt.k)
			for want := tt.expected; want != nil; got, want = got.Next, want.Next {
				if got == nil || got.Val != want.Val {
					t.Errorf("got %v, want %v", got, want)
				}
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := swapNodes2(cloneNode(tt.head), tt.k)
			for want := tt.expected; want != nil; got, want = got.Next, want.Next {
				if got == nil || got.Val != want.Val {
					t.Errorf("got %v, want %v", got, want)
				}
			}
		})
	}
}

func cloneNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	return &ListNode{
		Val:  head.Val,
		Next: cloneNode(head.Next),
	}
}
