package main

import "testing"

func TestValidPath(t *testing.T) {
	tests := []struct {
		name        string
		edges       [][]int
		n           int
		source      int
		destination int
		output      bool
	}{
		{"example1", [][]int{{0, 1}, {1, 2}, {2, 0}}, 3, 0, 2, true},
		{"example2", [][]int{{0, 7}, {0, 8}, {6, 1}, {2, 0}, {0, 4}, {5, 8}, {4, 7}, {1, 3}, {3, 5}, {6, 5}}, 10, 7, 5, true},
		{"example3", [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 6, 0, 5, false},
		{"example4", [][]int{}, 1, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := validPath(tt.edges, tt.n, tt.source, tt.destination)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := validPath2(tt.edges, tt.n, tt.source, tt.destination)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := validPath3(tt.edges, tt.n, tt.source, tt.destination)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}
}
