package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateRight(t *testing.T) {
	tests := []struct {
		name   string
		head   *ListNode
		k      int
		output *ListNode
	}{
		{
			name:   "example1",
			head:   &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
			k:      2,
			output: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}}},
		},
		{
			name:   "example2",
			head:   &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}},
			k:      4,
			output: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}}},
		},
		{
			name:   "example3",
			head:   &ListNode{Val: 1, Next: nil},
			k:      0,
			output: &ListNode{Val: 1, Next: nil},
		},
		{
			name:   "example4",
			head:   &ListNode{},
			k:      1,
			output: &ListNode{},
		},
		{
			name:   "example5",
			head:   &ListNode{Val: 1, Next: nil},
			k:      1,
			output: &ListNode{Val: 1, Next: nil},
		},
		{
			name:   "example6",
			head:   &ListNode{Val: 1, Next: nil},
			k:      99,
			output: &ListNode{Val: 1, Next: nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := rotateRight(tt.head, tt.k)
			assert.True(t, reflect.DeepEqual(res, tt.output))
		})
	}
}
