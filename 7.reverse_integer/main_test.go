package main

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		expected int
	}{{
		name:     "example1",
		x:        123,
		expected: 321,
	}, {
		name:     "example2",
		x:        -123,
		expected: -321,
	}, {
		name:     "example3",
		x:        120,
		expected: 21,
	}, {
		name:     "example4",
		x:        -910,
		expected: -19,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverse(tt.x)
			if got != tt.expected {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
