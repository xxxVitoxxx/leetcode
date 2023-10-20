package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name   string
		nums1  []int
		m      int
		nums2  []int
		n      int
		output []int
	}{
		{
			name:   "example1",
			nums1:  []int{1, 2, 3, 0, 0, 0},
			m:      3,
			nums2:  []int{2, 5, 6},
			n:      3,
			output: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:   "example2",
			nums1:  []int{1},
			m:      1,
			nums2:  []int{},
			n:      0,
			output: []int{1},
		},
		{
			name:   "example3",
			nums1:  []int{-10, 3, 5, 5, 6, 8, 9, 9, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			m:      9,
			nums2:  []int{-21, 1, 1, 5, 5, 7, 7, 9, 9, 11},
			n:      10,
			output: []int{-21, -10, 1, 1, 3, 5, 5, 5, 5, 6, 7, 7, 8, 9, 9, 9, 9, 10, 11},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)
			assert.Equal(t, tt.output, tt.nums1)
		})
	}
}

func TestMerge2(t *testing.T) {
	tests := []struct {
		name   string
		nums1  []int
		m      int
		nums2  []int
		n      int
		output []int
	}{
		{
			name:   "example1",
			nums1:  []int{1, 2, 3, 0, 0, 0},
			m:      3,
			nums2:  []int{2, 5, 6},
			n:      3,
			output: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:   "example2",
			nums1:  []int{1},
			m:      1,
			nums2:  []int{},
			n:      0,
			output: []int{1},
		},
		{
			name:   "example3",
			nums1:  []int{-10, 3, 5, 5, 6, 8, 9, 9, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			m:      9,
			nums2:  []int{-21, 1, 1, 5, 5, 7, 7, 9, 9, 11},
			n:      10,
			output: []int{-21, -10, 1, 1, 3, 5, 5, 5, 5, 6, 7, 7, 8, 9, 9, 9, 9, 10, 11},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge2(tt.nums1, tt.m, tt.nums2, tt.n)
			assert.Equal(t, tt.output, tt.nums1)
		})
	}
}

func TestMerge3(t *testing.T) {
	tests := []struct {
		name   string
		nums1  []int
		m      int
		nums2  []int
		n      int
		output []int
	}{
		{
			name:   "example1",
			nums1:  []int{1, 2, 3, 0, 0, 0},
			m:      3,
			nums2:  []int{2, 5, 6},
			n:      3,
			output: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:   "example2",
			nums1:  []int{1},
			m:      1,
			nums2:  []int{},
			n:      0,
			output: []int{1},
		},
		{
			name:   "example3",
			nums1:  []int{-10, 3, 5, 5, 6, 8, 9, 9, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			m:      9,
			nums2:  []int{-21, 1, 1, 5, 5, 7, 7, 9, 9, 11},
			n:      10,
			output: []int{-21, -10, 1, 1, 3, 5, 5, 5, 5, 6, 7, 7, 8, 9, 9, 9, 9, 10, 11},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge3(tt.nums1, tt.m, tt.nums2, tt.n)
			assert.Equal(t, tt.output, tt.nums1)
		})
	}
}

func TestMerge4(t *testing.T) {
	tests := []struct {
		name   string
		nums1  []int
		m      int
		nums2  []int
		n      int
		output []int
	}{
		{
			name:   "example1",
			nums1:  []int{1, 2, 3, 0, 0, 0},
			m:      3,
			nums2:  []int{2, 5, 6},
			n:      3,
			output: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:   "example2",
			nums1:  []int{1},
			m:      1,
			nums2:  []int{},
			n:      0,
			output: []int{1},
		},
		{
			name:   "example3",
			nums1:  []int{-10, 3, 5, 5, 6, 8, 9, 9, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			m:      9,
			nums2:  []int{-21, 1, 1, 5, 5, 7, 7, 9, 9, 11},
			n:      10,
			output: []int{-21, -10, 1, 1, 3, 5, 5, 5, 5, 6, 7, 7, 8, 9, 9, 9, 9, 10, 11},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge4(tt.nums1, tt.m, tt.nums2, tt.n)
			assert.Equal(t, tt.output, tt.nums1)
		})
	}
}
