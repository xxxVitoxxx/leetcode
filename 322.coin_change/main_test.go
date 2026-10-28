package main

import "testing"

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name     string
		coins    []int
		amount   int
		expected int
	}{{
		name:     "example1",
		coins:    []int{1, 2, 5},
		amount:   11,
		expected: 3,
	}, {
		name:     "example2",
		coins:    []int{2},
		amount:   3,
		expected: -1,
	}, {
		name:     "example3",
		coins:    []int{1},
		amount:   0,
		expected: 0,
	}, {
		name:     "example4",
		coins:    []int{2, 5, 10, 1},
		amount:   27,
		expected: 4,
	}, {
		name:     "example5",
		coins:    []int{186, 419, 83, 408},
		amount:   6249,
		expected: 20,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := coinChange(tt.coins, tt.amount)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := coinChange2(tt.coins, tt.amount)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
