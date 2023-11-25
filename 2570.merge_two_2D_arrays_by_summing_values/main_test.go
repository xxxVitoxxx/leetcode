package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeArrays(t *testing.T) {
	tests := []struct {
		name   string
		nums1  [][]int
		nums2  [][]int
		output [][]int
	}{
		{
			name:   "example1",
			nums1:  [][]int{{1, 2}, {2, 3}, {4, 5}},
			nums2:  [][]int{{1, 4}, {3, 2}, {4, 1}},
			output: [][]int{{1, 6}, {2, 3}, {3, 2}, {4, 6}},
		},
		{
			name:   "example2",
			nums1:  [][]int{{2, 4}, {3, 6}, {5, 5}},
			nums2:  [][]int{{1, 3}, {4, 3}},
			output: [][]int{{1, 3}, {2, 4}, {3, 6}, {4, 3}, {5, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := mergeArrays(tt.nums1, tt.nums2)
			assert.True(t, reflect.DeepEqual(res, tt.output))
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := mergeArrays2(tt.nums1, tt.nums2)
			assert.True(t, reflect.DeepEqual(res, tt.output))
		})
	}
}
