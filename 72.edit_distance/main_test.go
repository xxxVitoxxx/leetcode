package main

import "testing"

func TestMinDistance(t *testing.T) {
	tests := []struct {
		name     string
		word1    string
		word2    string
		expected int
	}{{
		name:     "example1",
		word1:    "horse",
		word2:    "ros",
		expected: 3,
	}, {
		name:     "example2",
		word1:    "intention",
		word2:    "execution",
		expected: 5,
	}, {
		name:     "example3",
		word1:    "carry",
		word2:    "terry",
		expected: 2,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minDistance(tt.word1, tt.word2)
			if got != tt.expected {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minDistance2(tt.word1, tt.word2)
			if got != tt.expected {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
