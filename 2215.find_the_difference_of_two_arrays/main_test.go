package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestFindDifference(t *testing.T) {
	tests := []struct {
		name   string
		nums1  []int
		nums2  []int
		output [][]int
	}{
		{"example1", []int{1, 2, 3}, []int{2, 4, 6}, [][]int{{1, 3}, {4, 6}}},
		{"example2", []int{1, 2, 3, 3}, []int{1, 1, 2, 2}, [][]int{{3}, {}}},
		{"example3", []int{-68, -80, -19, -94, 82, 21, -43}, []int{-63}, [][]int{{-68, -80, -19, -94, 82, 21, -43}, {-63}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findDifference(tt.nums1, tt.nums2)
			for i := 0; i < len(res); i++ {
				if !slices.EqualFunc(res[i], tt.output[i], func(v1, v2 int) bool {
					if v1 == v2 {
						return true
					}

					for _, num := range tt.output[i] {
						if v1 == num {
							return true
						}
					}

					return false
				}) {
					t.Errorf("got %v, want %v", res, tt.output)
				}
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findDifference2(tt.nums1, tt.nums2)
			for i := 0; i < len(res); i++ {
				if !slices.EqualFunc(res[i], tt.output[i], func(v1, v2 int) bool {
					if v1 == v2 {
						return true
					}

					for _, num := range tt.output[i] {
						if v1 == num {
							return true
						}
					}

					return false
				}) {
					t.Errorf("got %v, want %v", res, tt.output)
				}
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findDifference3(tt.nums1, tt.nums2)
			for i := 0; i < len(res); i++ {
				if !slices.EqualFunc(res[i], tt.output[i], func(v1, v2 int) bool {
					if v1 == v2 {
						return true
					}

					for _, num := range tt.output[i] {
						if v1 == num {
							return true
						}
					}

					return false
				}) {
					t.Errorf("got %v, want %v", res, tt.output)
				}
			}
		})
	}
}
