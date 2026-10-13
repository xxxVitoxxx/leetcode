package main

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	tests := []struct {
		name     string
		cost     []int
		expected int
	}{{
		name:     "exampl1",
		cost:     []int{10, 15, 20},
		expected: 15,
	}, {
		name:     "example2",
		cost:     []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
		expected: 6,
	}, {
		name:     "example3",
		cost:     []int{10, 15},
		expected: 10,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := minCostClimbingStairs(tt.cost)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := minCostClimbingStairs2(tt.cost)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := minCostClimbingStairs3(tt.cost)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
