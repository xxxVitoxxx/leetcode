package main

import "testing"

func TestGuessNumber(t *testing.T) {
	tests := []struct {
		name string
		pick int
		n    int
	}{
		{"example1", 6, 10},
		{"example2", 1, 1},
		{"example3", 1, 2},
		{"example4", 511, 100000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pick = tt.pick
			res := guessNumber(tt.n)
			if res != pick {
				t.Errorf("got %v, want %v", res, pick)
			}
		})
	}
}
