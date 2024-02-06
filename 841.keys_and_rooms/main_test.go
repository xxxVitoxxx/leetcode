package main

import "testing"

func TestCanVisitAllRooms(t *testing.T) {
	tests := []struct {
		name   string
		input  [][]int
		output bool
	}{
		{"example1", [][]int{{1}, {2}, {3}, {}}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := canVisitAllRooms(tt.input)
			if res != tt.output {
				t.Errorf("got: %v, want: %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := canVisitAllRooms2(tt.input)
			if res != tt.output {
				t.Errorf("got: %v, want: %v", res, tt.output)
			}
		})
	}
}
