package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestFindPeakElement(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		output []int
	}{
		{"example1", []int{1, 2, 3, 1}, []int{2}},
		{"example2", []int{1, 2, 1, 3, 5, 6, 4}, []int{1, 5}},
		{"example3", []int{1, 2, 3, 4, 5, 6}, []int{5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findPeakElement(tt.nums)
			if !slices.Contains(tt.output, res) {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findPeakElement2(tt.nums)
			if !slices.Contains(tt.output, res) {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findPeakElement3(tt.nums)
			if !slices.Contains(tt.output, res) {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}
}
