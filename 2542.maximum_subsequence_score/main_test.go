package main

import (
	"testing"
)

func TestMaxScore(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		k     int
		want  int64
	}{
		{
			name:  "example1",
			nums1: []int{1, 3, 3, 2},
			nums2: []int{2, 1, 3, 4},
			k:     3,
			want:  12,
		},
		{
			name:  "example2",
			nums1: []int{4, 2, 3, 1, 1},
			nums2: []int{7, 5, 10, 9, 6},
			k:     1,
			want:  30,
		},
		{
			name:  "example3",
			nums1: []int{23, 16, 20, 7, 3},
			nums2: []int{19, 21, 22, 22, 12},
			k:     3,
			want:  1121,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxScore(tt.nums1, tt.nums2, tt.k)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
