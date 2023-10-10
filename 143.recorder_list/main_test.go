package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReorderList(t *testing.T) {
	tests := []struct {
		name   string
		input  *ListNode
		output []int
	}{
		{
			name: "example1",
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			output: []int{1, 4, 2, 3},
		},
		{
			name: "example2",
			input: &ListNode{
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
			output: []int{1, 5, 2, 4, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.input)
			var res []int
			for tt.input != nil {
				res = append(res, tt.input.Val)
				tt.input = tt.input.Next
			}

			assert.Len(t, res, len(tt.output))
			for i, output := range tt.output {
				assert.Equal(t, output, res[i])
			}
		})
	}
}
