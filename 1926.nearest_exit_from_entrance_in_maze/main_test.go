package main

import "testing"

func TestNearestExit(t *testing.T) {
	tests := []struct {
		name     string
		maze     [][]byte
		entrance []int
		output   int
	}{
		{
			"example1",
			[][]byte{
				{'+', '.', '+', '+', '+', '+', '+'},
				{'+', '.', '+', '.', '.', '.', '+'},
				{'+', '.', '+', '.', '+', '.', '+'},
				{'+', '.', '.', '.', '.', '.', '+'},
				{'+', '+', '+', '+', '.', '+', '.'},
			},
			[]int{0, 1},
			7,
		},
		{
			"example2",
			[][]byte{
				{'+', '+', '.', '+'},
				{'.', '.', '.', '+'},
				{'+', '+', '+', '.'},
			},
			[]int{1, 2},
			1,
		},
		{
			"example3",
			[][]byte{
				{'+', '+', '+'},
				{'.', '.', '.'},
				{'+', '+', '+'},
			},
			[]int{1, 0},
			2,
		},
		{
			"example4",
			[][]byte{
				{'+', '.', '+', '+', '+', '+', '+'},
				{'+', '.', '+', '.', '.', '.', '+'},
				{'+', '.', '+', '.', '+', '.', '+'},
				{'+', '.', '.', '.', '.', '.', '+'},
				{'+', '+', '+', '+', '.', '+', '.'},
			},
			[]int{0, 1},
			7,
		},
		{
			"example5",
			[][]byte{
				{'.', '+'},
			},
			[]int{0, 0},
			-1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := nearestExit(tt.maze, tt.entrance)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}
}
