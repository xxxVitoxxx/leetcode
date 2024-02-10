package main

import "testing"

func TestFindCircleNum(t *testing.T) {
	tests := []struct {
		name   string
		input  [][]int
		output int
	}{
		{"example1", [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}, 2},
		{"example2", [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findCircleNum(tt.input)
			if res != tt.output {
				t.Errorf("got: %v, want: %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findCircleNum2(tt.input)
			if res != tt.output {
				t.Errorf("got: %v, want: %v", res, tt.output)
			}
		})
	}
}
