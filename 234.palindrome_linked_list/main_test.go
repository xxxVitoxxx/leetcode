package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name   string
		input  *ListNode
		output bool
	}{
		{
			name: "example1",
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
			output: true,
		},
		{
			name: "example2",
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
			output: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isPalindrome(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isPalindrome2(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
