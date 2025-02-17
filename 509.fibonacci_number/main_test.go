package main

import "testing"

func TestFib(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{{
		name: "example1",
		n:    2,
		want: 1,
	}, {
		name: "example2",
		n:    3,
		want: 2,
	}, {
		name: "example3",
		n:    4,
		want: 3,
	}, {
		name: "example4",
		n:    30,
		want: 832040,
	}}

	for _, tt := range tests {
		t.Run("approach1 "+tt.name, func(t *testing.T) {
			got := fib(tt.n)
			if got != tt.want {
				t.Errorf("got %v, want: %v", got, tt.want)
			}
		})

		t.Run("approach2 "+tt.name, func(t *testing.T) {
			got := fib2(tt.n)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})

		t.Run("approach3 "+tt.name, func(t *testing.T) {
			got := fib3(tt.n)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})

		t.Run("approach4 "+tt.name, func(t *testing.T) {
			got := fib4(tt.n)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
