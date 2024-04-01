package main

import (
	"testing"
)

func TestCloseStrings(t *testing.T) {
	tests := []struct {
		name   string
		word1  string
		word2  string
		output bool
	}{
		{"example1", "abc", "bca", true},
		{"example2", "a", "aa", false},
		{"example3", "cabbba", "abbccc", true},
		{"example4", "cabba", "abccc", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := closeStrings(tt.word1, tt.word2)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := closeStrings2(tt.word1, tt.word2)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := closeStrings3(tt.word1, tt.word2)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}
}
