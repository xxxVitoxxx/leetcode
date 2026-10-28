package main

import "testing"

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]byte
		expected int
	}{{
		name:     "example1",
		grid:     [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}},
		expected: 1,
	}, {
		name:     "example2",
		grid:     [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}},
		expected: 3,
	}, {
		name:     "example3",
		grid:     [][]byte{{'1', '1', '0'}, {'1', '0', '1'}, {'0', '0', '1'}},
		expected: 2,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := numIslands(copyGrid(tt.grid))
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := numIslands2(copyGrid(tt.grid))
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := numIslands3(copyGrid(tt.grid))
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := numIslands4(copyGrid(tt.grid))
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := numIslands5(copyGrid(tt.grid))
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}

func copyGrid(grid [][]byte) [][]byte {
	newGrid := make([][]byte, len(grid))
	for i := range grid {
		newGrid[i] = make([]byte, len(grid[i]))
		copy(newGrid[i], grid[i])
	}
	return newGrid
}
