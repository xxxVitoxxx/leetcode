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
		{
			"example2",
			[][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			-1,
		},
		{
			"example3",
			[][]int{
				{0, 2},
			},
			0,
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

func TestOrangesRotting2(t *testing.T) {
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
		{
			"example2",
			[][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			-1,
		},
		{
			"example3",
			[][]int{
				{0, 2},
			},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := orangesRotting2(tt.grid)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}
}
