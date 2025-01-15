package main

import (
	"reflect"
	"testing"
)

func TestZigzagTraversal(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want []int
	}{{
		name: "example1",
		grid: [][]int{{1, 2}, {3, 4}},
		want: []int{1, 4},
	}, {
		name: "example2",
		grid: [][]int{{2, 1}, {2, 1}, {2, 1}},
		want: []int{2, 1, 2},
	}, {
		name: "example3",
		grid: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		want: []int{1, 3, 5, 7, 9},
	}, {
		name: "",
		grid: [][]int{{1, 2, 1, 3}, {5, 15, 7, 3}, {10, 4, 14, 12}},
		want: []int{1, 1, 3, 15, 10, 14},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := zigzagTraversal(tt.grid)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := zigzagTraversal2(tt.grid)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := zigzagTraversal3(tt.grid)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
