package main

import "testing"

func TestShortestPathBinaryMatrix(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{{
		name:     "example1",
		grid:     [][]int{{0, 1}, {1, 0}},
		expected: 2,
	}, {
		name:     "example2",
		grid:     [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}},
		expected: 4,
	}, {
		name:     "example3",
		grid:     [][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}},
		expected: -1,
	}, {
		name:     "example4",
		grid:     [][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}},
		expected: -1,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := copyGrid(tt.grid)
			actual := shortestPathBinaryMatrix(grid)
			if actual != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, actual)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := copyGrid(tt.grid)
			actual := shortestPathBinaryMatrix2(grid)
			if actual != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, actual)
			}
		})
	}
}

func copyGrid(grid [][]int) [][]int {
	newGrid := make([][]int, len(grid))
	for i := range grid {
		newGrid[i] = make([]int, len(grid[i]))
		copy(newGrid[i], grid[i])
	}
	return newGrid
}
