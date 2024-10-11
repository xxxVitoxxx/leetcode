package main

import "testing"

func TestMissingNumber(t *testing.T) {
	for _, tt := range []struct {
		nums []int
		want int
	}{{
		nums: []int{3, 0, 1},
		want: 2,
	}, {
		nums: []int{0, 1},
		want: 2,
	}, {
		nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
		want: 8,
	}, {
		nums: []int{7, 3, 4, 6, 1, 5, 0},
		want: 2,
	}} {
		nums1 := make([]int, len(tt.nums))
		copy(nums1, tt.nums)
		got := missingNumber(nums1)
		if got != tt.want {
			t.Errorf("got: %d, want: %d", got, tt.want)
		}

		nums2 := make([]int, len(tt.nums))
		copy(nums2, tt.nums)
		got = missingNumber2(nums2)
		if got != tt.want {
			t.Errorf("got: %d, want: %d", got, tt.want)
		}

		nums3 := make([]int, len(tt.nums))
		copy(nums3, tt.nums)
		got = missingNumber3(nums3)
		if got != tt.want {
			t.Errorf("got: %d, want: %d", got, tt.want)
		}

		nums4 := make([]int, len(tt.nums))
		copy(nums4, tt.nums)
		got = missingNumber4(nums4)
		if got != tt.want {
			t.Errorf("got: %d, want: %d", got, tt.want)
		}
	}
}
