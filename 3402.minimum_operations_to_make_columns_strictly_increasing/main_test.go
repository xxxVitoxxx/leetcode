package main

import "testing"

func TestMinimumOperations(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{{
		name: "example1",
		grid: [][]int{{3, 2}, {1, 3}, {3, 4}, {0, 1}},
		want: 15,
	}, {
		name: "example2",
		grid: [][]int{{3, 2, 1}, {2, 1, 0}, {1, 2, 3}},
		want: 12,
	}, {
		name: "example3",
		grid: [][]int{{0, 0}, {0, 0}},
		want: 2,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumOperations(tt.grid)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
