package main

import "testing"

func TestChange(t *testing.T) {
	tests := []struct {
		name     string
		amount   int
		coins    []int
		expected int
	}{{
		name:     "example1",
		amount:   5,
		coins:    []int{1, 2, 5},
		expected: 4,
	}, {
		name:     "example2",
		amount:   3,
		coins:    []int{2},
		expected: 0,
	}, {
		name:     "example3",
		amount:   10,
		coins:    []int{10},
		expected: 1,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := change(tt.amount, tt.coins)
			if got != tt.expected {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
