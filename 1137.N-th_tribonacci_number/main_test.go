package main

import "testing"

func TestTribonacci(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "example1",
			n:    4,
			want: 4,
		},
		{
			name: "example2",
			n:    25,
			want: 1389537,
		},
		{
			name: "example3",
			n:    37,
			want: 2082876103,
		},
		{
			name: "example4",
			n:    0,
			want: 0,
		},
		{
			name: "example5",
			n:    1,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tribonacci(tt.n); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tribonacci2(tt.n); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
