package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name   string
		l1     *ListNode
		l2     *ListNode
		answer *ListNode
	}{
		{
			name: "example1",
			l1: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			l2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			answer: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 8,
					},
				},
			},
		},
		{
			name: "example2",
			l1: &ListNode{
				Val: 0,
			},
			l2: &ListNode{
				Val: 0,
			},
			answer: &ListNode{
				Val: 0,
			},
		},
		{
			name: "example3",
			l1: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
								Next: &ListNode{
									Val: 9,
									Next: &ListNode{
										Val: 9,
									},
								},
							},
						},
					},
				},
			},
			l2: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
					},
				},
			},
			answer: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 0,
									Next: &ListNode{
										Val: 0,
										Next: &ListNode{
											Val: 1,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := addTwoNumbers(tt.l1, tt.l2)
			assert.EqualValues(t, tt.answer, res)
		})
	}
}
