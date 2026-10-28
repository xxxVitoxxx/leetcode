package main

import "testing"

func TestExist(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		word     string
		expected bool
	}{{
		name:     "example1",
		board:    [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
		word:     "ABCCED",
		expected: true,
	}, {
		name:     "example2",
		board:    [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
		word:     "SEE",
		expected: true,
	}, {
		name:     "example3",
		board:    [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
		word:     "ABCB",
		expected: false,
	}, {
		name:     "example4",
		board:    [][]byte{{'a', 'b'}},
		word:     "ba",
		expected: true,
	}, {
		name:     "example5",
		board:    [][]byte{{'a', 'a'}},
		word:     "aaa",
		expected: false,
	}, {
		name:     "example6",
		board:    [][]byte{{'a', 'b'}, {'c', 'd'}},
		word:     "acdb",
		expected: true,
	}, {
		name:     "example7",
		board:    [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}},
		word:     "ABCESEEEFS",
		expected: true,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := exist(tt.board, tt.word)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
