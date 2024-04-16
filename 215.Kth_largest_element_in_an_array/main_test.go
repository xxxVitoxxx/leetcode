package main

import (
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		k      int
		output int
	}{
		{"example1", []int{3, 2, 1, 5, 6, 4}, 2, 5},
		{"example2", []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		{"example3", []int{3, 1, 2, 4}, 2, 3},
		{"example4", []int{2, 1}, 2, 1},
		{"example5", []int{1, 2, 3, 4, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 2, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			res := findKthLargest(nums, tt.k)
			if res != tt.output {
				t.Errorf("got %v want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			res := findKthLargest2(nums, tt.k)
			if res != tt.output {
				t.Errorf("got %v want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			res := findKthLargest3(nums, tt.k)
			if res != tt.output {
				t.Errorf("got %v want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			res := findKthLargest4(nums, tt.k)
			if res != tt.output {
				t.Errorf("got %v want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			res := findKthLargest5(nums, tt.k)
			if res != tt.output {
				t.Errorf("got %v want %v", res, tt.output)
			}
		})
	}
}
