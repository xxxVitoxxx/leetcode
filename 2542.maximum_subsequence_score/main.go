package main

import (
	"container/heap"
	"sort"
)

// heap
// time complexity: O(n * log n)
//   - we need to sort nums2, it takes O(n * log n).
//   - then we iterate over pairs of length n. at each iteration step i,
//     we remove the smallest element from minHeap and add one element pairs[i][0] to it.
//     both the inserting and removing operations to priority queue of size k take O(log k) time.
//   - to sum up, the overall time complexity is O(n * log n), because k <= n.
//
// space complexity: O(n)
//   - we store every pair [nums1[i], nums2[i]] in a 2-d slice pairs, it takes O(n) space.
//   - the in-place sorting method also uses some additional space,
//     in Golang, it use pdqsort and it use O(log n) space.
//   - the priority queue contains at most k elements thus it takes O(k) space.
func maxScore(nums1 []int, nums2 []int, k int) int64 {
	pairs := make([][]int, len(nums1))
	for i := 0; i < len(pairs); i++ {
		pairs[i] = []int{nums1[i], nums2[i]}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] > pairs[j][1]
	})

	minHeap := &MinHeap{}
	sum := 0
	for i := 0; i < k; i++ {
		sum += pairs[i][0]
		heap.Push(minHeap, pairs[i][0])
	}

	answer := sum * pairs[k-1][1]

	for i := k; i < len(nums1); i++ {
		smallest := heap.Pop(minHeap)
		sum += pairs[i][0] - smallest.(int)

		heap.Push(minHeap, pairs[i][0])
		answer = max(answer, sum*pairs[i][1])
	}

	return int64(answer)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

type MinHeap []int

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
