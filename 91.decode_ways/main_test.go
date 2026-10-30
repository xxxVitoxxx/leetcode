package main

import "testing"

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{{
		name:     "example1",
		s:        "12",
		expected: 2,
	}, {
		name:     "example2",
		s:        "226",
		expected: 3,
	}, {
		name:     "example3",
		s:        "06",
		expected: 0,
	}, {
		name:     "example4",
		s:        "11106",
		expected: 2,
	}, {
		name:     "example5",
		s:        "110",
		expected: 1,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := numDecodings(tt.s)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := numDecodings2(tt.s)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
