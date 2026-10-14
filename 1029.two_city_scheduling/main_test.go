package main

import "testing"

func TestTwoCitySchedCost(t *testing.T) {
	tests := []struct {
		name     string
		costs    [][]int
		expected int
	}{{
		name:     "example1",
		costs:    [][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}},
		expected: 110,
	}, {
		name:     "example2",
		costs:    [][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}},
		expected: 1859,
	}, {
		name:     "example3",
		costs:    [][]int{{515, 563}, {451, 713}, {537, 709}, {343, 819}, {855, 779}, {457, 60}, {650, 359}, {631, 42}},
		expected: 3086,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := twoCitySchedCost(tt.costs)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
