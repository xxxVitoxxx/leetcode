package main

import "testing"

func TestFindLonelyPixel(t *testing.T) {
	tests := []struct {
		name     string
		picture  [][]byte
		expected int
	}{{
		name:     "example1",
		picture:  [][]byte{{'W', 'W', 'B'}, {'W', 'B', 'W'}, {'B', 'W', 'W'}},
		expected: 3,
	}, {
		name:     "example2",
		picture:  [][]byte{{'B', 'B', 'B'}, {'B', 'B', 'W'}, {'B', 'B', 'B'}},
		expected: 0,
	}, {
		name: "example3",
		picture: [][]byte{{'B', 'W', 'W', 'W', 'W', 'B', 'W', 'B', 'B', 'W'},
			{'B', 'B', 'B', 'W', 'W', 'W', 'B', 'W', 'B', 'W'},
			{'B', 'B', 'B', 'B', 'W', 'W', 'W', 'B', 'W', 'W'},
			{'B', 'W', 'W', 'B', 'W', 'B', 'B', 'W', 'W', 'W'},
			{'W', 'W', 'B', 'W', 'B', 'B', 'B', 'W', 'B', 'B'},
			{'W', 'B', 'B', 'W', 'W', 'W', 'B', 'W', 'W', 'W'},
			{'B', 'W', 'W', 'B', 'B', 'B', 'W', 'W', 'W', 'W'},
			{'W', 'W', 'W', 'B', 'B', 'B', 'B', 'W', 'W', 'W'},
			{'W', 'W', 'B', 'B', 'W', 'W', 'W', 'W', 'B', 'W'},
			{'W', 'W', 'W', 'B', 'B', 'B', 'W', 'W', 'W', 'B'}},
		expected: 0,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := findLonelyPixel(tt.picture)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := findLonelyPixel2(tt.picture)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
