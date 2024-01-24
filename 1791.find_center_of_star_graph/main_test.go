package main

import "testing"

func TestFindCenter(t *testing.T) {
	tests := []struct {
		name   string
		edges  [][]int
		output int
	}{
		{"example1", [][]int{{1, 2}, {2, 3}, {4, 2}}, 2},
		{"example2", [][]int{{1, 1}, {5, 1}, {1, 3}, {1, 4}}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findCenter(tt.edges)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}
}
