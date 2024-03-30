package main

import (
	"testing"
)

func TestEqualPairs(t *testing.T) {
	tests := []struct {
		name   string
		grid   [][]int
		output int
	}{
		{"example1", [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}, 1},
		{"example2", [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := equalPairs(tt.grid)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := equalPairs2(tt.grid)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := equalPairs3(tt.grid)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}
}
