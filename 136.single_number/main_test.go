package main

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		output int
	}{
		{"example1", []int{2, 2, 1}, 1},
		{"example2", []int{4, 1, 2, 1, 2}, 4},
		{"example3", []int{1}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := singleNumber(tt.nums)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := singleNumber2(tt.nums)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := singleNumber3(tt.nums)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}
}
