package main

import "testing"

func TestCountSubstrings(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{{
		name:     "example1",
		s:        "abc",
		expected: 3,
	}, {
		name:     "example2",
		s:        "aaa",
		expected: 6,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := countSubstrings(tt.s)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := countSubstrings2(tt.s)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := countSubstrings3(tt.s)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
