package main

import "testing"

func TestMinMeetingRooms(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      int
	}{
		{
			name:      "example1",
			intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
			want:      2,
		},
		{
			name:      "example2",
			intervals: [][]int{{7, 10}, {2, 4}},
			want:      1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minMeetingRooms(tt.intervals)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
