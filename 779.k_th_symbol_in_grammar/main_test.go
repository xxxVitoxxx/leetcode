package main

import "testing"

func TestKthGrammer(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want int
	}{{
		name: "example 1",
		n:    1,
		k:    1,
		want: 0,
	}, {
		name: "example 2",
		n:    2,
		k:    1,
		want: 0,
	}, {
		name: "example 3",
		n:    2,
		k:    2,
		want: 1,
	}, {
		name: "example 4",
		n:    3,
		k:    4,
		want: 0,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := kthGrammar(test.n, test.k); got != test.want {
				t.Errorf("got: %v, want: %v", got, test.want)
			}
		})
	}
}
