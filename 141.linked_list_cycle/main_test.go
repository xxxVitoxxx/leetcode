package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasCycle(t *testing.T) {
	e1l4 := &ListNode{Val: -4}
	e1l3 := &ListNode{Val: 0, Next: e1l4}
	e1l2 := &ListNode{Val: 2, Next: e1l3}
	e1l1 := &ListNode{Val: 3, Next: e1l2}
	e1l4.Next = e1l2
	/*
	   3 -> 2 -> 0 -> -4
	        ↑          |
	         - - - - - -
	*/

	e2l2 := &ListNode{Val: 2}
	e2l1 := &ListNode{Val: 1, Next: e2l2}
	e2l2.Next = e2l2
	/*
	   1 -> 2
	   ↑    |
	   - - -
	*/

	e3l1 := &ListNode{Val: 1}
	/*
	   1
	*/

	tests := []struct {
		name   string
		input  *ListNode
		output bool
	}{
		{
			name:   "example1",
			input:  e1l1,
			output: true,
		},
		{
			name:   "example2",
			input:  e2l1,
			output: true,
		},
		{
			name:   "example3",
			input:  e3l1,
			output: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := hasCycle(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
