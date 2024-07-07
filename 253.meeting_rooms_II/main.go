package main

import (
	"container/heap"
	"sort"
)

// heap(priority queue)
// time complexity: O(n * log n)
//   - there are two major portions that take up time here. one is sorting of the array that
//     takes O(n * log n) considering that the array consists of n elements.
//   - then we have the min-heap. in the worst case, all n meetings will collide with each other.
//     in any case we have n add operations on the heap. in the worst case we will have
//     n extract-min operations as well. overall complexity being O(n * log n) since
//     extract-min operation on a heap takes O(log n)
//
// space complexity: O(n)
//   - because we construct the min-heap and that can contain n elements in the worst case as
//     described above in the time complexity section. hence, the space complexity is O(n).
func minMeetingRooms(intervals [][]int) int {
	if len(intervals) <= 1 {
		return len(intervals)
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	minHeap := &MinHeap{}
	heap.Push(minHeap, intervals[0][1])

	for i := 1; i < len(intervals); i++ {
		if (*minHeap)[0] <= intervals[i][0] {
			heap.Pop(minHeap)
		}
		heap.Push(minHeap, intervals[i][1])
	}

	return minHeap.Len()
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

// chronological ordering
// time complexity: O(n * log n)
// space complexity: O(n)
func minMeetingRooms2(intervals [][]int) int {
	if len(intervals) <= 1 {
		return len(intervals)
	}

	// separate out the start and end time
	startTimes, endTimes := make([]int, len(intervals)), make([]int, len(intervals))
	for i := 0; i < len(intervals); i++ {
		startTimes[i] = intervals[i][0]
		endTimes[i] = intervals[i][1]
	}

	// sort start and end time in increasing order
	sort.Slice(startTimes, func(i, j int) bool {
		return startTimes[i] < startTimes[j]
	})
	sort.Slice(endTimes, func(i, j int) bool {
		return endTimes[i] < endTimes[j]
	})

	var endPoint int
	for startPoint := 0; startPoint < len(startTimes); startPoint++ {
		// if there is a meeting that has ended by the time the meeting at 'startPoint' startTimes
		if startTimes[startPoint] >= endTimes[endPoint] {
			endPoint++
		}
	}

	return len(endTimes) - endPoint
}
