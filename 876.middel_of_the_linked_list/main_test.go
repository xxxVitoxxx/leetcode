package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddleNode(t *testing.T) {
	e1l5 := &ListNode{Val: 5}
	e1l4 := &ListNode{Val: 4, Next: e1l5}
	e1l3 := &ListNode{Val: 3, Next: e1l4}
	e1l2 := &ListNode{Val: 2, Next: e1l3}
	e1l1 := &ListNode{Val: 1, Next: e1l2}

	e2l6 := &ListNode{Val: 6}
	e2l5 := &ListNode{Val: 5, Next: e2l6}
	e2l4 := &ListNode{Val: 4, Next: e2l5}
	e2l3 := &ListNode{Val: 3, Next: e2l4}
	e2l2 := &ListNode{Val: 2, Next: e2l3}
	e2l1 := &ListNode{Val: 1, Next: e2l2}

	tests := []struct {
		name   string
		input  *ListNode
		output *ListNode
	}{
		{"example1", e1l1, e1l3},
		{"example2", e2l1, e2l4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := middleNode(tt.input)
			assert.True(t, reflect.DeepEqual(*res, *tt.output))
		})
	}
}
