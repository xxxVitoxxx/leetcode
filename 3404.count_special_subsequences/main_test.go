package main

import "testing"

func TestNumberOfSubsequences(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int64
	}{{
		name: "example1",
		nums: []int{1, 2, 3, 4, 3, 6, 1},
		want: 1,
	}, {
		name: "example2",
		nums: []int{3, 4, 3, 4, 3, 4, 3, 4},
		want: 3,
	}, {
		name: "example3",
		nums: []int{43, 376, 716, 10, 102, 330, 800, 176, 914, 858, 774, 806, 26, 26, 10, 10, 10, 43, 555, 972, 43, 43, 10, 43, 10, 653, 26, 10, 10, 26, 43, 10, 43, 10, 10, 43, 26, 10, 26, 10, 10, 10, 26, 43, 26, 10, 10, 10, 26, 26, 43, 10, 26, 43, 43, 43, 26, 43, 10, 10, 43, 43, 43, 43, 10, 42, 43, 42, 10, 10, 26, 43, 26, 43, 10, 43, 26, 26, 43, 10, 26, 10, 43, 10, 43, 43, 450, 43, 10, 26, 10, 43, 26, 10, 26, 43, 10, 664, 26, 26, 26, 6, 10, 43, 43, 6, 43, 26, 10, 10, 10, 43, 10, 839, 10, 10, 43, 26, 307, 10, 43, 43, 10, 6, 26, 43, 10, 10, 424, 43, 10, 26, 243, 26, 10, 43, 43, 26, 43, 26, 43, 26, 10, 26, 10, 812, 6, 6, 43, 26, 352, 26, 26, 10, 338, 43, 394, 43, 43, 850, 6, 26, 43, 10, 10, 43, 43, 26, 70, 43, 26, 43, 43, 26, 26, 43, 43, 26, 10, 26, 43, 26, 26, 43, 43, 536, 43, 43, 26, 26, 43, 43, 43, 10, 26, 26, 26, 26, 26, 26, 10, 26, 43, 26, 26, 26, 26, 26, 10, 43, 26, 43, 674, 26, 42, 43, 10, 10, 10, 70, 6, 43, 43, 43, 10, 6, 10, 43, 6, 10, 10, 43, 26, 43, 757, 10, 10, 26, 42, 26, 10, 43, 26, 10, 42, 43, 10, 315, 42, 26, 26, 43, 465, 26, 26, 26, 10, 43, 43, 26, 43, 43, 26, 631, 43, 26, 43, 10, 133, 10, 10, 26, 10, 10, 43, 42, 43, 10, 26, 26, 10, 26, 10, 10, 10, 10, 10, 10, 6, 70, 6, 43, 26, 43, 465, 43, 10, 43, 10, 26, 26, 10, 26, 10, 26, 43, 43, 10, 70, 10, 6, 43, 26, 26, 70, 70, 70, 43, 26, 6, 42, 10, 388, 26, 26, 26, 26, 6, 10, 43, 43, 10, 70, 10, 10, 70, 43, 70, 43, 43, 482, 70, 10, 6, 10, 26, 10, 43, 26, 70, 70, 10, 26, 10, 70, 10, 375, 26, 847, 10, 205, 10, 43, 43, 563, 10, 10, 795, 6, 43, 43, 10, 10, 26, 26, 10, 10, 631, 505, 10, 43, 10, 10, 26, 43, 70, 287, 6, 42, 26, 70, 26, 10, 10, 10, 10, 43, 6, 70, 42, 947, 121, 390, 26, 47, 10, 43, 70, 43, 10, 10, 26, 10, 43, 10, 26, 70, 954, 43, 719, 772, 26, 43, 43, 26, 43, 812, 43, 26, 10, 43, 43, 6, 10, 26, 26, 43, 43, 43, 6, 10, 26, 503, 43, 42, 26, 26, 43, 151, 43, 43, 288, 43, 10, 43, 989, 26, 26, 6, 43, 26, 47, 10, 43, 26, 26, 42, 70, 10, 10, 10, 26, 26, 6, 70, 26, 70, 70, 6, 26, 10, 70, 43, 26, 983, 20, 10, 10, 43, 26, 43, 43, 26, 43, 10, 10, 26, 10, 70, 26, 26, 26, 43, 42, 26, 26, 26, 42, 10, 70, 26, 70, 10, 43, 26, 451, 10, 861, 607, 6, 10, 864, 26, 43, 42, 6, 42, 70, 26, 26, 47, 43, 26, 43, 43, 10, 26, 47, 10, 6},
		want: 219130937,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSubsequences(tt.nums); got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
