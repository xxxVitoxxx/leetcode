package main

import "testing"

func TestMinReorder(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		connections [][]int
		output      int
	}{
		{"example1", 6, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := minReorder(tt.n, tt.connections)
			if res != tt.output {
				t.Fatalf("expected: %v, got: %v", tt.output, res)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := minReorder2(tt.n, tt.connections)
			if res != tt.output {
				t.Fatalf("expected: %v, got: %v", tt.output, res)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := minReorder3(tt.n, tt.connections)
			if res != tt.output {
				t.Fatalf("expected: %v, got: %v", tt.output, res)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := minReorder4(tt.n, tt.connections)
			if res != tt.output {
				t.Fatalf("expected: %v, got: %v", tt.output, res)
			}
		})
	}
}
