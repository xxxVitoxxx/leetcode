package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		name   string
		nums1  []int
		nums2  []int
		output []int
	}{
		{"example1", []int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{"example2", []int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
		{"example3", []int{1, 1, 1, 1}, []int{9, 1, 1}, []int{1, 1}},
		{"example4", []int{9, 2, 2, 6, 6, 1, 2, 9}, []int{2, 8, 5, 2, 6}, []int{2, 2, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := intersect(tt.nums1, tt.nums2)
			assert.ElementsMatch(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := intersect2(tt.nums1, tt.nums2)
			assert.ElementsMatch(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := intersect3(tt.nums1, tt.nums2)
			assert.ElementsMatch(t, tt.output, res)
		})
	}
}

func TestIntersect4(t *testing.T) {
	tests := []struct {
		name   string
		nums1  []int
		nums2  []int
		output []int
	}{
		{"example1", []int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{"example2", []int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
		{"example3", []int{1, 1, 1, 1}, []int{9, 1, 1}, []int{1, 1}},
		{"example4", []int{9, 2, 2, 6, 6, 1, 2, 9}, []int{2, 8, 5, 2, 6}, []int{2, 2, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := intersect4(tt.nums1, tt.nums2)
			assert.ElementsMatch(t, tt.output, res)
		})
	}
}
