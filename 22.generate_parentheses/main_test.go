package main

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []string
	}{{
		name:     "example1",
		n:        3,
		expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
	}, {
		name:     "example2",
		n:        1,
		expected: []string{"()"},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateParenthesis(tt.n)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
