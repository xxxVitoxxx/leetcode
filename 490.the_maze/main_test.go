package main

import "testing"

func TestHasPath(t *testing.T) {
	tests := []struct {
		name        string
		maze        [][]int
		start       []int
		destination []int
		want        bool
	}{
		{
			name: "example1",
			maze: [][]int{
				{0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0},
				{1, 1, 0, 1, 1},
				{0, 0, 0, 0, 0},
			},
			start:       []int{0, 4},
			destination: []int{4, 4},
			want:        true,
		},
		{
			name: "example2",
			maze: [][]int{
				{0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0},
				{1, 1, 0, 1, 1},
				{0, 0, 0, 0, 0},
			},
			start:       []int{0, 4},
			destination: []int{3, 2},
			want:        false,
		},
		{
			name: "example3",
			maze: [][]int{
				{0, 0, 0, 0, 0},
				{1, 1, 0, 0, 1},
				{0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1},
				{0, 1, 0, 0, 0},
			},
			start:       []int{4, 3},
			destination: []int{0, 1},
			want:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPath(tt.maze, tt.start, tt.destination); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPath2(tt.maze, tt.start, tt.destination); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
