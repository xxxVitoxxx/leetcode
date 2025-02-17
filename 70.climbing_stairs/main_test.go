package main

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{{
		name: "example1",
		n:    1,
		want: 1,
	}, {
		name: "example2",
		n:    2,
		want: 2,
	}, {
		name: "example3",
		n:    3,
		want: 3,
	}, {
		name: "example4",
		n:    6,
		want: 13,
	}, {
		name: "example5",
		n:    5,
		want: 8,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbStairs(tt.n)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			got := climbStairs2(tt.n)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			got := climbStairs3(tt.n)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			got := climbStairs4(tt.n)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
