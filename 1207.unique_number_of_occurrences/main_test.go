package main

import (
	"testing"
)

func TestUniqueOccurrences(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		output bool
	}{
		{"example1", []int{1, 2, 2, 1, 1, 3}, true},
		{"example2", []int{1, 2}, false},
		{"example3", []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := uniqueOccurrences(tt.arr)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := uniqueOccurrences2(tt.arr)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}
}
