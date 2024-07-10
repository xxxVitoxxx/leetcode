package main

import (
	"testing"

	"github.com/xxxVitoxxx/leetcode/testfunc"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "example1",
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
		{
			name: "example2",
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
		{
			name: "example3",
			nums: []int{2, 7, 3, 4, 4, 4, 2, 1, 4, 4, 2},
			k:    2,
			want: []int{2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.nums, tt.k)
			if !testfunc.ElementsMatch(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}

		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent2(tt.nums, tt.k)
			if !testfunc.ElementsMatch(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent3(tt.nums, tt.k)
			if !testfunc.ElementsMatch(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
