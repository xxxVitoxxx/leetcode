package main

import "testing"

func TestTotalNQueens(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "example1",
			n:    4,
			want: 2,
		},
		{
			name: "example2",
			n:    1,
			want: 1,
		},
		{
			name: "example3",
			n:    8,
			want: 92,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := totalNQueens(tt.n)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := totalNQueens2(tt.n)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
