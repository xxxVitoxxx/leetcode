package main

import "testing"

func TestGetSum(t *testing.T) {
	for _, tt := range []struct {
		a    int
		b    int
		want int
	}{{
		a:    -1,
		b:    -9,
		want: -10,
	}, {
		a:    -10,
		b:    5,
		want: -5,
	}, {
		a:    1,
		b:    2,
		want: 3,
	}, {
		a:    2,
		b:    3,
		want: 5,
	}} {
		got := getSum(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("got: %d, want: %d", got, tt.want)
		}

		got = getSum2(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("got: %d, want: %d", got, tt.want)
		}
	}
}
