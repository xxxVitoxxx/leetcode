package main

import "testing"

func TestDetectSquares(t *testing.T) {
	ds := Constructor()
	ds.Add([]int{3, 10})
	ds.Add([]int{11, 2})
	ds.Add([]int{3, 2})

	if result := ds.Count([]int{11, 10}); result != 1 {
		t.Errorf("got %v, want %v", result, 1)
	}

	if result := ds.Count([]int{14, 8}); result != 0 {
		t.Errorf("got %v, want %v", result, 0)
	}

	ds.Add([]int{11, 2})

	if result := ds.Count([]int{11, 10}); result != 2 {
		t.Errorf("got %v, want %v", result, 2)
	}
}
