package main

import "testing"

func TestAddBinary(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected string
	}{{
		name:     "example1",
		a:        "11",
		b:        "1",
		expected: "100",
	}, {
		name:     "example2",
		a:        "1010",
		b:        "1011",
		expected: "10101",
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addBinary(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("got %v, want: %v", got, tt.expected)
			}
		})
	}
}
