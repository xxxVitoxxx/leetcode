package main

import "testing"

func TestCanAttendMeetings(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  bool
	}{{
		name:      "example1",
		intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
		expected:  false,
	}, {
		name:      "example1",
		intervals: [][]int{{7, 10}, {2, 4}},
		expected:  true,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := canAttendMeetings(tt.intervals)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
