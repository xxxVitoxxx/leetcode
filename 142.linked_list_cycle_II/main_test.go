package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectCycle(t *testing.T) {
	e1l4 := &ListNode{Val: 4}
	e1l3 := &ListNode{Val: 0, Next: e1l4}
	e1l2 := &ListNode{Val: 2, Next: e1l3}
	e1l1 := &ListNode{Val: 3, Next: e1l2}
	// e1l2 is the cycle begin
	e1l4.Next = e1l2

	e2l2 := &ListNode{Val: 2}
	e2l1 := &ListNode{Val: 1, Next: e2l2}
	// e2l1 is the cycle begin
	e2l2.Next = e2l1

	// no cycle
	e3l1 := &ListNode{Val: 1}

	tests := []struct {
		name       string
		input      *ListNode
		cycleBegin *ListNode
	}{
		{"example1", e1l1, e1l2},
		{"example2", e2l1, e2l1},
		{"example3", e3l1, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cycleBegin := detectCycle(tt.input)
			if cycleBegin != nil {
				assert.Equal(t, tt.cycleBegin.Val, cycleBegin.Val)
			} else {
				assert.Nil(t, tt.cycleBegin)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cycleBegin := detectCycle2(tt.input)
			if cycleBegin != nil {
				assert.Equal(t, tt.cycleBegin.Val, cycleBegin.Val)
			} else {
				assert.Nil(t, tt.cycleBegin)
			}
		})
	}
}
