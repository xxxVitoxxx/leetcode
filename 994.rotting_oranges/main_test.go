package main

import "testing"

func TestOrangesRotting(t *testing.T) {
	tests := []struct {
		name   string
		grid   [][]int
		output int
	}{
		{
			"example1",
			[][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := orangesRotting(tt.grid)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}
}
