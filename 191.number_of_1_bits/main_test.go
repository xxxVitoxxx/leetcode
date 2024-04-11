package main

import (
	"testing"
)

func TestHammingWeight(t *testing.T) {
	tests := []struct {
		name   string
		n      int
		output int
	}{
		{"example1", 11, 3},
		{"example2", 128, 1},
		{"example3", 2147483645, 30},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := hammingWeight(tt.n)
			if res != tt.output {
				t.Errorf("got %v want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := hammingWeight2(tt.n)
			if res != tt.output {
				t.Errorf("got %v want %v", res, tt.output)
			}
		})
	}
}
